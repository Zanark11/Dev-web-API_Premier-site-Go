package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	// Route de la page d'accueil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// On vérifie que le visiteur demande bien la page d'accueil
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		// On envoie Bonjour au navigateur
		fmt.Fprintln(w, "Bonjour")
	})

	// Route pour afficher la liste des artistes
	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {

		// On prépare notre liste d'artistes
		artistes := []string{
			"Koba la D",
			"Leto",
			"Beendo Z",
		}

		// On charge le fichier HTML
		tmpl, err := template.ParseFiles("templates/artists.html")

		if err != nil {
			http.Error(w, "Erreur lors du chargement de la page", http.StatusInternalServerError)
			return
		}

		// On transmet les artistes au template
		tmpl.Execute(w, artistes)
	})

	// Route pour afficher la fiche d'un artiste
	http.HandleFunc("/artists/", func(w http.ResponseWriter, r *http.Request) {

		// On récupère le chemin demandé
		chemin := r.URL.Path

		// On retire /artists/ pour récupérer l'identifiant
		idTexte := strings.TrimPrefix(chemin, "/artists/")

		// On transforme le texte en nombre entier
		id, err := strconv.Atoi(idTexte)

		// On vérifie si l'identifiant est valide
		if err != nil {
			http.Error(w, "Identifiant invalide", http.StatusBadRequest)
			return
		}

		// On prépare le nom de l'artiste
		var artiste string

		if id == 1 {
			artiste = "Koba la D"
		} else if id == 2 {
			artiste = "Leto"
		} else if id == 3 {
			artiste = "Beendo Z"
		} else {
			http.Error(w, "Artiste introuvable", http.StatusNotFound)
			return
		}

		// On charge le template de la fiche artiste
		tmpl, err := template.ParseFiles("templates/artist.html")

		if err != nil {
			http.Error(w, "Erreur lors du chargement de la page", http.StatusInternalServerError)
			return
		}

		// On transmet le nom de l'artiste au template
		tmpl.Execute(w, artiste)
	})

	// On permet au navigateur d'accéder aux fichiers CSS
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// On démarre le serveur sur le port 8080
	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
