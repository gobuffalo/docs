<% seoDescription("Comment écrire des requêtes brutes avec Pop") %>
<% seoKeywords(["buffalo", "go", "golang", "bdd", "brut", "requête", "custom"]) %>

<%= h1("Requêtes brutes") %>

Parfois vous aurez besoin d'écrire des requêtes à la main au lieu de laisser Pop les générer pour vous. Dans ce chapitre, vous allez apprendre comment écrire des requêtes SQL brutes en utilisant Pop.

## Écrire une requête brute

### Select

```go
player := Player{}
q := db.RawQuery("SELECT * FROM players WHERE id = ?", 1)
err := q.Find(&player, id)
```

### Update

```go
err := db.RawQuery("UPDATE players SET instrument = ? WHERE id = ?", "guitar", 1).Exec()
```

### Delete

```go
err := db.RawQuery("DELETE FROM players WHERE id = ?", 1).Exec()
```

## Syntaxe des jetons de valeurs

Avec `RawQuery`, vous pouvez continuer à utiliser les jetons `?` pour sécuriser les valeurs d'entrée. Il n'est pas nécessaire d'utiliser la syntaxe de votre base de données.