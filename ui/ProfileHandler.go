package ui

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type ProfileHandler struct {
	Template *template.Template
}

func NewProfileHandler() *ProfileHandler {
	return &ProfileHandler{Template: Tmpl}
}

func (h *ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProfileHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	idStr := fields[len(fields)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if id < 1 || id > len(IndexVM.Artists) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	vm := IndexVM.Artists[id-1]
	err = h.Template.ExecuteTemplate(w, "Profile", vm)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
