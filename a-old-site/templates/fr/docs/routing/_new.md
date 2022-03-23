## Créer une nouvelle application Buffalo (et son routeur)

La configuration de l'app se trouve dans le fichier `app.go`.

```go
a := buffalo.New(buffalo.Options{
  Env:         ENV,
  SessionName: "_coke_session",
})
```

La configuration par défaut devrait satisfaire la plupart de vos besoins, mais vous êtes libre de la modifier pour mieux y répondre.

La liste des options est disponible ici : [https://godoc.org/github.com/gobuffalo/buffalo#Options](https://godoc.org/github.com/gobuffalo/buffalo#Options)
