package http

import (
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func serve_file(w http.ResponseWriter, req *http.Request) {
	web_dir, err := fs.Sub(fsys, "web")
	if err != nil {
		log.Fatal(err)
	}
	http.FileServer(http.FS(web_dir)).ServeHTTP(w, req)
}

func trimBaseURL(baseURL string) string {
	if baseURL == "/" {
		return "/"
	}
	return strings.TrimRight(baseURL, "/")
}

// Router takes a base URL for prefixing all URLs returned
// to the client (so you don't need to burn a subdomain for
// this service)
// and returns an http.Handler that can accept all traffic
func Router(baseURL string) *chi.Mux {
	rV := chi.NewRouter()
	rV.Use(middleware.Logger)
	rV.Use(middleware.RealIP)
	baseUrlWithoutLastSlash := trimBaseURL(baseURL)
	rV.Route(baseUrlWithoutLastSlash, func(r chi.Router) {
		r.Handle("/*", http.StripPrefix(baseURL, http.HandlerFunc(serve_file)))
	})
	return rV
}
