package main

import (
	"log"

	"github.com/Vaelatern/gokrazy-statuspage/internal/app"
)

func main() {
	//r := myhttp.Router("/")
	//r.Use(middleware.Logger)
	//r.Use(middleware.RealIP)
	//http.ListenAndServe(":3000", r)
	log.Fatal(app.Entrypoint())
}
