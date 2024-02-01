package main

import (
	"log"
	"matthewhope/groupie-tracker/router"
	"matthewhope/groupie-tracker/ui"
	"net/http"
	"regexp"
)

func main() {
	mux := http.NewServeMux()
	serveStaticFiles(mux)
	r := router.New()
	r.AddHandler(regexp.MustCompile(`^/$`), ui.NewIndexHandler())
	r.AddHandler(regexp.MustCompile(`^/profile/[0-9]+$`), ui.NewProfileHandler())
	mux.Handle("/", r)
	err := http.ListenAndServe(":1234", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func serveStaticFiles(mux *http.ServeMux) {
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}
