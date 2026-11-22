# Présentation

Ce projet consiste à créer un premier site web en utilisant le langage Go.

Le serveur utilise le package `net/http` pour gérer les requêtes HTTP, `html/template` pour générer les pages HTML et des fichiers CSS pour la mise en forme du site.

## Installation et lancement

### Prérequis

- Go installé sur votre ordinateur.

### Lancer le projet

Cloner le dépôt GitHub :

```bash
git clone https://github.com/Zanark11/Dev-web-API_Premier-site-Go.git
```

Entrer dans le dossier :

```bash
cd Dev-web-API_Premier-site-Go
```

Lancer le serveur :

```bash
go run .
```

Le site est accessible à l'adresse :

http://localhost:8080

## Routes disponibles

| Route | Description |
|---|---|
| `/` | Page d'accueil |
| `/artists` | Liste des artistes |
| `/artists/1` | Fiche de Koba la D |
| `/artists/2` | Fiche de Leto |
| `/artists/3` | Fiche de Beendo Z |

## Gestion des erreurs HTTP

- `400 Bad Request` : identifiant invalide.
- `404 Not Found` : artiste introuvable.
- `200 OK` : requête réussie.

## Technologies utilisées

- Go
- HTML
- CSS
- Git et GitHub
