package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// Page contient les données passées aux templates.
type Page struct {
	Title  string // Titre de l'onglet
	Active string // Page active pour le menu (home, about, portfolio, contact)
}

var templates map[string]*template.Template

// loadTemplates compile chaque page avec le layout de base.
func loadTemplates() {
	templates = make(map[string]*template.Template)
	pages := []string{"home", "about", "portfolio", "contact"}
	for _, p := range pages {
		t := template.Must(template.ParseFiles(
			"templates/base.html",
			"templates/"+p+".html",
		))
		templates[p] = t
	}
}

func render(w http.ResponseWriter, name string, data Page) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates[name].ExecuteTemplate(w, "base", data); err != nil {
		log.Printf("erreur template %s: %v", name, err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
	}
}

func main() {
	loadTemplates()

	mux := http.NewServeMux()

	// Fichiers statiques (CSS, JS, images)
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		render(w, "home", Page{Title: "Francky Relex Ndacayisaba — Portfolio", Active: "home"})
	})
	mux.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		render(w, "about", Page{Title: "À propos — Francky Relex Ndacayisaba", Active: "about"})
	})
	mux.HandleFunc("GET /portfolio", func(w http.ResponseWriter, r *http.Request) {
		render(w, "portfolio", Page{Title: "Portfolio — Francky Relex Ndacayisaba", Active: "portfolio"})
	})
	mux.HandleFunc("GET /contact", func(w http.ResponseWriter, r *http.Request) {
		render(w, "contact", Page{Title: "Contact — Francky Relex Ndacayisaba", Active: "contact"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("✦ Portfolio en ligne sur http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
