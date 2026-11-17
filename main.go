package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Artist struct {
	ID   int
	Name string
}

var artists = []Artist{
	{ID: 1, Name: "Eminem"},
	{ID: 2, Name: "Angèle"},
	{ID: 3, Name: "Stromae"},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func afficherListeArtistes(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/list.html", "template/layout.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	q := r.URL.Query().Get("q")
	if q == "" {
		t.Execute(w, artists)
	} else {
		nouvelleListe := []Artist{}
		for _, a := range artists {
			if strings.HasPrefix(strings.ToLower(a.Name), strings.ToLower(q)) {
				nouvelleListe = append(nouvelleListe, a)
			}
		}
		t.Execute(w, nouvelleListe)
	}
}

func afficherArtiste(w http.ResponseWriter, r *http.Request) {
	idTexte := r.PathValue("id")
	id, err := strconv.Atoi(idTexte)
	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}
	for _, a := range artists {
		if a.ID == id {
			t, err := template.ParseFiles("template/artiste.html", "template/layout.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			t.Execute(w, a)
			return
		}
	}
	http.Error(w, "artiste introuvable", http.StatusNotFound)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/artists", afficherListeArtistes)
	mux.HandleFunc("/artists/{id}", afficherArtiste)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", mux)
}
