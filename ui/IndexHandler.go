package ui

import (
	"html/template"
	"log"
	"net/http"
)

type IndexHandler struct {
	Template *template.Template
}

func NewIndexHandler() *IndexHandler {
	tmpl, err := template.ParseGlob("./templates/*.go.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	return &IndexHandler{Template: tmpl}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *IndexHandler) get(w http.ResponseWriter, r *http.Request) {
	err := h.Template.ExecuteTemplate(w, "Index", IndexVM)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
