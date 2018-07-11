<% seoDescription("Cookies") %>
<% seoKeywords(["buffalo", "go", "golang", "http", "cookie"]) %>

<%= h1("Cookies") %>

Un cookie HTTP est un petit morceau de données qu'un serveur envoie au navigateur web d'un utilisateur. Le navigateur peut stocker cette donnée et la renvoyer au même serveur, même après redémarrage du navigateur (contrairement à [une session](/en/docs/sessions)).

Les cookies (HTTP) sont souvent utilisés pour sauvegarder l'état de l'utlisateur courant (comme par example pour savoir si l'utilisateur est connecté). Consultez [https://golang.org/pkg/net/http/#Cookie](https://golang.org/pkg/net/http/#Cookie) pour plus d'informations sur la manière de gérer les cookies en Go.

<%= title("Créer un cookie") %>

```go
func MyHandler(c buffalo.Context) error {
  // ...
  c.Cookies().Set("user_id", user.ID, 30 * 24 * time.Hour)
  // ...
}
```

<%= title("Créer un cookie périssable") %>

```go
func MyHandler(c buffalo.Context) error {
  // ...
  exp := time.Now().Add(365 * 24 * time.Hour) // expire dans 1 an
  c.Cookies().SetWithExpirationTime("user_id", user.ID, exp)
  // ...
}
```

<%= title("Récupérer un cookie") %>

```go
func MyHandler(c buffalo.Context) error {
  value, err := c.Cookies().Get("user_id")
  if err != nil {
    return err
  }
  return c.Render(200, r.String(value))
}
```

<%= title("Supprimer un cookie") %>


```go
func MyHandler(c buffalo.Context) error {
  c.Cookies().Delete("user_id")
  // ...
}
```
