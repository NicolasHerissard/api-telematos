# CLI-API

Une API REST simple construite avec [Gin](https://github.com/gin-gonic/gin) en Go. Cette API permet de gérer des utilisateurs, des produits, des rapports et leurs relations.

## Présentation

Ce projet est une API REST qui expose plusieurs endpoints pour interagir avec une base de données. Les fonctionnalités incluent :
- La récupération des utilisateurs.
- La récupération des produits.
- La génération de rapports.
- La gestion des relations entre produits et utilisateurs.
- La création d'utilisateurs.

L'API est conçue pour être extensible et facile à utiliser.

---

## Installation

### Prérequis

- [Go](https://golang.org/dl/) installé sur votre machine.
- Une base de données configurée et accessible.

### Étapes d'installation

1. Clonez ce dépôt :

   ```bash
   git clone <URL_DU_DEPOT>
   cd cli-api
   ```

2. Installez les dépendances :

   ```bash
   go mod tidy
   ```

3. Compilez l'application :

   ```bash
   go build -o api.exe
   ```

4. Exécutez l'application :

   ```bash
   go run [main.go](http://_vscodecontentref_/0) -u <utilisateur> -p <mot_de_passe> -d <nom_base_de_donnees>
   ```

5. Testez l'application :

   ```bash
   curl http://localhost:4444/users
   ```

## Endpoints

### GET `/users`

Récupère la liste des utilisateurs.

#### Exemple de réponse

```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
]
```

### GET `/products`

Récupère la liste des produits.

#### Exemple de réponse

```json
[
  {
    "id": 1,
    "name": "Product A",
    "price": 100
  }
]
```

## Structure du projet
    ```bash
cli-api/
├── [api.exe]
├── [go.mod]
├── [go.sum]
├── [main.go]
└── [model.go]
    ```

- **`main.go`** : Point d'entrée principal de l'application.
- **`model.go`** : Contient les fonctions pour interagir avec la base de données.