<%= h1("Error Handling") %>

Une `error` est la manière que Go utilise pour prévenir que quelque chose s'est mal passé. Dans ce chapitre, vous allez apprendre comment retourner des erreurs depuis un contrôleur ; et comment Buffalo rattrape les erreurs non-traitées.

<%= title("Retourner une erreur depuis un contrôleur") %>

La façon la plus simple de renvoyer une erreur HTTP est de retourner une erreur Go standard :
```go
func MyHandler(c buffalo.Context) error {
  // Retourner une erreur go classique : cela provoquera un statut HTTP 500.
  return errors.New("boom!")
}
```

Une erreur `nil` produira une réponse HTTP 200 :
```go
func MyHandler(c buffalo.Context) error {
  // HTTP 200
  return nil
}
```

Si vous avez besoin de personnaliser l'erreur avec un message ou un code HTTP différent, utilisez la méthode <%= doclink("github.com/gobuffalo/buffalo#DefaultContext.Error") %> :
```go
func MyHandler(c buffalo.Context) error {
  // En utilisant la méthode Error du Context.
  // Dans cet exemple, le résultat sera une réponse HTTP 401.
  return c.Error(401, errors.New("Unauthorized!"))
}
```

<%= title("Rattrapage des erreurs par défaut (mode développement)") %>

En mode développement (`GO_ENV=development`), Buffalo génère des pages d'erreur détaillées pour vous permettre de débogguer facilement.

<figure>
  <img src="/assets/images/500_example.png" title="screenshot">
  <figcaption>Un exemple d'erreur `500` en mode développement.</figcaption>
</figure>

Si vous utilisez un `Content-Type` JSON ou XML, l'erreur sera formattée en respectant le type demandé :

```json
{
  "error": "could not find test/",
  "trace": "could not find test/\ngithub.com/gobuffalo/gobuffalo/vendor/github.com/gobuffalo/buffalo.(*App).fileServer.func1\n\t/home/michalakst/go/src/github.com/gobuffalo/gobuffalo/vendor/github.com/gobuffalo/buffalo/route_mappings.go:97\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1947\nnet/http.StripPrefix.func1\n\t/usr/local/go/src/net/http/server.go:1986\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1947\ngithub.com/gobuffalo/gobuffalo/vendor/github.com/gorilla/mux.(*Router).ServeHTTP\n\t/home/michalakst/go/src/github.com/gobuffalo/gobuffalo/vendor/github.com/gorilla/mux/mux.go:162\ngithub.com/gobuffalo/gobuffalo/vendor/github.com/markbates/refresh/refresh/web.ErrorChecker.func1\n\t/home/michalakst/go/src/github.com/gobuffalo/gobuffalo/vendor/github.com/markbates/refresh/refresh/web/web.go:23\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1947\ngithub.com/gobuffalo/gobuffalo/vendor/github.com/gobuffalo/buffalo.(*App).ServeHTTP\n\t/home/michalakst/go/src/github.com/gobuffalo/gobuffalo/vendor/github.com/gobuffalo/buffalo/server.go:127\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2694\nnet/http.(*conn).serve\n\t/usr/local/go/src/net/http/server.go:1830\nruntime.goexit\n\t/usr/local/go/src/runtime/asm_amd64.s:2361",
  "code": 404
}
```

```xml
&lt;response code="404">
  &lt;error>could not find test/&lt;/error>
  &lt;trace>could not find test/&#xA;github.com/gobuffalo/gobuffalo/vendor/github.com/gobuffalo/buffalo.(*App).fileServer.func1&#xA;&#x9;/home/michalakst/go/src/github.com/gobuffalo/gobuffalo/vendor/github.com/gobuffalo/buffalo/route_mappings.go:97&#xA;net/http.HandlerFunc.ServeHTTP&#xA;&#x9;/usr/local/go/src/net/http/server.go:1947&#xA;net/http.StripPrefix.func1&#xA;&#x9;/usr/local/go/src/net/http/server.go:1986&#xA;net/http.HandlerFunc.ServeHTTP&#xA;&#x9;/usr/local/go/src/net/http/server.go:1947&#xA;github.com/gobuffalo/gobuffalo/vendor/github.com/gorilla/mux.(*Router).ServeHTTP&#xA;&#x9;/home/michalakst/go/src/github.com/gobuffalo/gobuffalo/vendor/github.com/gorilla/mux/mux.go:162&#xA;github.com/gobuffalo/gobuffalo/vendor/github.com/markbates/refresh/refresh/web.ErrorChecker.func1&#xA;&#x9;/home/michalakst/go/src/github.com/gobuffalo/gobuffalo/vendor/github.com/markbates/refresh/refresh/web/web.go:23&#xA;net/http.HandlerFunc.ServeHTTP&#xA;&#x9;/usr/local/go/src/net/http/server.go:1947&#xA;github.com/gobuffalo/gobuffalo/vendor/github.com/gobuffalo/buffalo.(*App).ServeHTTP&#xA;&#x9;/home/michalakst/go/src/github.com/gobuffalo/gobuffalo/vendor/github.com/gobuffalo/buffalo/server.go:127&#xA;net/http.serverHandler.ServeHTTP&#xA;&#x9;/usr/local/go/src/net/http/server.go:2694&#xA;net/http.(*conn).serve&#xA;&#x9;/usr/local/go/src/net/http/server.go:1830&#xA;runtime.goexit&#xA;&#x9;/usr/local/go/src/runtime/asm_amd64.s:2361&lt;/trace>
&lt;/response>
```

En mode production (`GO_ENV=production`), Buffalo ne génère pas ces pages, vu qu'elles donneraient des informations très utiles à des personnes mal intentionnées. Des pages d'erreur neutres sont générées à la place.

<%= title("Gestion personnalisée des erreurs", {}) %>

Bien que Buffalo se charge de gérer les erreurs pour vous sans n'avoir rien à faire, il peut être utile de gérer les erreurs d'une autre façon. Vous pouvez pour cela associer des codes HTTP à des contrôleurs spécifiques. De cette manière, vous pouvez gérer ces erreurs comme bon vous semble.

```go
app = buffalo.New(buffalo.Options{
  Env: ENV,
})

// On associe à l'erreur HTTP 422 un traitement spécifique.
// Toutes les autres erreurs gardent le traitement par défaut.
app.ErrorHandlers[422] = func(status int, err error, c buffalo.Context) error {
  res := c.Response()
  res.WriteHeader(422)
  res.Write([]byte(fmt.Sprintf("Oops!! There was an error %s", err.Error())))
  return nil
}

app.GET("/oops", MyHandler)

func MyHandler(c buffalo.Context) error {
  return c.Error(422, errors.New("Oh no!"))
}
```

```text
GET /oops -> [422] Oh no!
```

Dans l'exemple ci-dessus, toute erreur de votre application renvoyant un statut `422` sera rattrapée par le contrôleur personnalisé, et renverra donc le message d'erreur `Oops!! There was an error` avec le texte de l'erreur.
