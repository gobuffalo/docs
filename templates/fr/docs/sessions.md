<% seoDescription("Sessions") %>
<% seoKeywords(["buffalo", "go", "golang", "http", "session"]) %>

<%= h1("Sessions") %>

Une session HTTP est un stockage de données non-persistant, détruit lors de la fermeture du navigateur web (dans une configuration classique). Ce stockage peut être utilisé pour conserver des messages flash, ou tout autre donnée temporaire propre à un utilisateur. Utilisez les [cookies](/fr/docs/cookies) à la place si vous avez besoin d'un stockage plus persistant côté utilisateur.

La session est directement disponible depuis le `buffalo.Context`, depuis un contrôleur.

```go
func MyHandler(c buffalo.Context) error {
  s := c.Session()
}
```

<%= partial("fr/docs/sessions/type.md") %>
<%= partial("fr/docs/sessions/store.md") %>
<%= partial("fr/docs/sessions/complex.md") %>
<%= partial("fr/docs/sessions/save.md") %>
<%= partial("fr/docs/sessions/null.md") %>

