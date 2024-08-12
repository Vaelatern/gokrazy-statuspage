package http

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Card struct {
	Title              string
	Metric             string
	IconText           string
	ShowHeartPulseIcon bool
	ShowServerIcon     bool
	IsOK               bool
	IsError            bool
}

type Payload struct {
	Cards []Card
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

	payload := Payload{Cards: []Card{
		Card{Title: "Self", Metric: "0ms", IsOK: true, ShowHeartPulseIcon: true},
		Card{Title: "Self", ShowServerIcon: true, IsOK: true},
		Card{Title: "Office1", Metric: "317ms", IsOK: true, ShowHeartPulseIcon: true},
		Card{Title: "Server", Metric: "68ms", IsOK: true, ShowHeartPulseIcon: true},
		Card{Title: "Server Responding", IsOK: true, IconText: "HTTP"},
		Card{Title: "Office2", Metric: "334ms", IsOK: true, ShowHeartPulseIcon: true},
		Card{Title: "Office Synology", Metric: "66ms", IsOK: true, ShowHeartPulseIcon: true},
		Card{Title: "Office Synology Responding", IsOK: true, IconText: "HTTP"},
		Card{Title: "Office3", IsError: true, ShowHeartPulseIcon: true},
		Card{Title: "Home Synology", Metric: "41ms", IsOK: true, ShowHeartPulseIcon: true},
		Card{Title: "Home Synology Responding", IsOK: true, IconText: "HTTP"}}}

	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Serving a template")
		// Load and parse the template file
		t, err := template.ParseFS(web_dir, tmpl, "definitions.tmpl")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}

		fmt.Println("Serving a template ", w)
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
		r.Handle("/*", http.StripPrefix(baseURL, http.HandlerFunc(serve_file)))
	})
	return rV
}
