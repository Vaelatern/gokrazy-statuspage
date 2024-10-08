package http

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"

	"github.com/Vaelatern/gokrazy-statuspage/internal/metrics"
)

type Payload struct {
	SecondsPoll int
	NumCols     int
	Cards       []metrics.Card
}

func serve_file(w http.ResponseWriter, req *http.Request) {
	web_dir, err := fs.Sub(fsys, "web")
	if err != nil {
		log.Fatal(err)
	}
	http.FileServer(http.FS(web_dir)).ServeHTTP(w, req)
}

func trim_base_url(baseURL string) string {
	if baseURL == "/" {
		return "/"
	}
	return strings.TrimRight(baseURL, "/")
}

func serve_template(tmpl string) func(http.ResponseWriter, *http.Request) {
	web_dir, err := fs.Sub(fsys, "web")
	if err != nil {
		log.Fatal(err)
	}

	return func(w http.ResponseWriter, req *http.Request) {
		payloadConfig := viper.Get("tests")
		if payloadConfig == nil {
			t, err := template.ParseFS(web_dir, "not-configured.html", "definitions.tmpl")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Println("Error parsing template:", err)
				return
			}
			err = t.Execute(w, nil)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Println("Error executing template:", err)
			}
			return
		}
		payload := Payload{NumCols: viper.GetInt("columns"),
			SecondsPoll: viper.GetInt("poll-frequency"),
			Cards:       metrics.AllCards(payloadConfig.([]interface{}))}
		// Load and parse the template file
		t, err := template.ParseFS(web_dir, tmpl, "definitions.tmpl")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}

		// Execute the template with data (if any)
		err = t.Execute(w, payload)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	}
}

// Router takes a base URL for prefixing all URLs returned
// to the client (so you don't need to burn a subdomain for
// this service)
// and returns an http.Handler that can accept all traffic
func Router(baseURL string) *chi.Mux {
	rV := chi.NewRouter()
	rV.Use(middleware.Logger)
	rV.Use(middleware.RealIP)
	baseUrlWithoutLastSlash := trim_base_url(baseURL)
	rV.Route(baseUrlWithoutLastSlash, func(r chi.Router) {
		r.Handle("/", http.StripPrefix(baseURL, http.HandlerFunc(serve_template("index.tmpl"))))
		r.Handle("/tmpl/main.html", http.StripPrefix(baseURL, http.HandlerFunc(serve_template("tmpl/main.html"))))
		r.Handle("/*", http.StripPrefix(baseURL, http.HandlerFunc(serve_file)))
	})
	return rV
}
