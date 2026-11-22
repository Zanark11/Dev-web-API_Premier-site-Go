package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// Structure qui représente un album
type Album struct {
	Nom   string
	Photo string
}

// Structure qui représente une certification
type Certification struct {
	Album  string
	Niveau string
}

// Structure qui représente un artiste
type Artiste struct {
	Nom            string
	Photo          string
	Albums         []Album
	Certifications []Certification
}

func main() {

	// On prépare les informations de nos artistes
	artistes := []Artiste{

		{
			Nom:   "Koba la D",
			Photo: "/static/images/koba-la.jpg",
			Albums: []Album{
				{
					Nom:   "L'Affranchi",
					Photo: "/static/images/koba-album1.jpg",
				},
				{
					Nom:   "VII",
					Photo: "/static/images/koba-album2.jpg",
				},
				{
					Nom:   "Détail",
					Photo: "/static/images/koba-album3.jpg",
				},
			},
			Certifications: []Certification{
				{
					Album:  "Album 1",
					Niveau: "Disque d'or",
				},
				{
					Album:  "Album 3",
					Niveau: "Disque de platine",
				},
				{
					Album:  "Album 1",
					Niveau: "Double disque de platine",
				},
			},
		},

		{
			Nom:   "Leto",
			Photo: "/static/images/leto.jpg",
			Albums: []Album{
				{
					Nom:   "THUG CEREMONY",
					Photo: "/static/images/leto-album1.jpg",
				},
				{
					Nom:   "Jusq'aux étoiles",
					Photo: "/static/images/leto-album2.jpg",
				},
				{
					Nom:   "LIFE",
					Photo: "/static/images/leto-album3.jpg",
				},
			},
			Certifications: []Certification{
				{
					Album:  "Album 1",
					Niveau: "Disque d'or",
				},
				{
					Album:  "Album 1",
					Niveau: "Disque de platine",
				},
				{
					Album:  "Album 0",
					Niveau: "Disque de diamant",
				},
			},
		},

		{
			Nom:   "Beendo Z",
			Photo: "/static/images/beendo-z.png",
			Albums: []Album{
				{
					Nom:   "L'Elu",
					Photo: "/static/images/beendo-album1.jpg",
				},
				{
					Nom:   "De la fontaine",
					Photo: "/static/images/beendo-album2.jpg",
				},
				{
					Nom:   "PROPA",
					Photo: "/static/images/beendo-album3.jpg",
				},
			},
			Certifications: []Certification{
				{
					Album:  "Album 1",
					Niveau: "Disque d'or",
				},
				{
					Album:  "Album 0",
					Niveau: "Disque de platine",
				},
				{
					Album:  "Album 0",
					Niveau: "Disque de diamant",
				},
			},
		},
	}

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

		// On vérifie que l'identifiant correspond à un artiste
		if id < 1 || id > len(artistes) {
			http.Error(w, "Artiste introuvable", http.StatusNotFound)
			return
		}

		// On récupère l'artiste correspondant à l'identifiant
		artiste := artistes[id-1]

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
