## Stockage des sessions

Par défaut, Buffalo stocke les sessions via un cookie en utilisant [`sessions.CookieStore`](http://www.gorillatoolkit.org/pkg/sessions#CookieStore).

Vous pouvez remplacer ce système de stockage en configurant votre application via l'option `SessionStore` :

```go
app = buffalo.New(buffalo.Options{
  Env:         ENV,
  SessionName: "_coke_session",
  SessionStore: sessions.NewCookieStore([]byte("some session secret")),
})
```

La variable d'environnement `SESSION_SECRET` doit être configurée avant de démarrer l'application. Si ce n'est pas le cas, vous verrez un avertissement dans vos logs disant que votre session n'est pas sécurisée.

Pour plus d'informations sur ce sujet, consultez la documentation de [`buffalo.Options`](https://godoc.org/github.com/gobuffalo/buffalo#Options).

