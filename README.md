# TP : mon premier serveur en Go

Petit site web écrit en Go, sans dépendance externe. Il affiche une liste d'artistes et la fiche de chacun. Les pages HTML sont générées avec `html/template`.

## Prérequis

- **Go 1.22 ou plus récent** (nécessaire pour les routes du type `/artists/{id}` et `r.PathValue`)

Pour vérifier la version installée :

```bash
go version
```

## Lancer le serveur

Depuis la racine du projet (le dossier qui contient `main.go`) :

```bash
go run .
```

Le serveur écoute sur le port **8080**. Ouvrir ensuite [http://localhost:8080](http://localhost:8080) dans un navigateur.

> Le serveur doit être lancé depuis la racine du projet, sinon il ne trouve pas les dossiers `template/` et `static/`.

Pour arrêter le serveur : `Ctrl + C`.

## Routes disponibles

| Méthode | Route | Description |
|---|---|---|
| GET | `/` | Page d'accueil (affiche « Hello ») |
| GET | `/artists` | Liste de tous les artistes |
| GET | `/artists?q=...` | Recherche : artistes dont le nom commence par le texte donné (sans tenir compte des majuscules) |
| GET | `/artists/{id}` | Fiche de l'artiste correspondant à l'id |
| GET | `/static/...` | Fichiers statiques (CSS) |

### Codes de réponse

| Cas | Code |
|---|---|
| Page trouvée | `200 OK` |
| Id qui n'est pas un nombre (ex. `/artists/abc`) | `400 Bad Request` |
| Id inconnu (ex. `/artists/999`) | `404 Not Found` |
| Route inexistante | `404 Not Found` |
| Template introuvable ou invalide | `500 Internal Server Error` |

## Exemples

Avec `curl` (sous PowerShell, utiliser `curl.exe`) :

```bash
curl -i http://localhost:8080/
curl -i http://localhost:8080/artists
curl -i "http://localhost:8080/artists?q=em"
curl -i http://localhost:8080/artists/1
curl -i http://localhost:8080/artists/999   # 404
curl -i http://localhost:8080/artists/abc   # 400
```

L'option `-i` affiche le statut et les en-têtes de la réponse (dont le `Content-Type`).

## Structure du projet

```
.
├── go.mod
├── main.go              # serveur, routes, handlers et données
├── template/
│   ├── layout.html      # morceaux communs : en-tête ("header") et pied de page ("footer")
│   ├── list.html        # page liste des artistes
│   └── artiste.html     # page fiche d'un artiste
└── static/
    └── style.css        # feuille de style
```

## Données

Les artistes sont écrits en dur dans `main.go`, dans la variable `artists`. Pour en ajouter un, il suffit d'ajouter une ligne dans cette liste avec un `ID` unique.