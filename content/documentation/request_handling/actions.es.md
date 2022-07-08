---
name: Controlador de Acciones
seoDescription: "How to define and use Action Controllers?"
seoKeywords: ["buffalo", "go", "golang", "actions", "controller", "generator"]
weight: 2
aliases:
  - /docs/actions
  - /es/docs/actions
---

# Controlador de Acciones

En este capítulo, aprenderás cómo funcionan los controladores de acciones; y cómo puedes generarlos usando los generadores integrados.

## Qué es un controlador?

Los controladores son la *C* del [modelo MVC](https://es.wikipedia.org/wiki/Modelo%E2%80%93vista%E2%80%93controlador). Estos manejan la lógica dada la desición del enrutador, y generan una respuesta apropiada.

Por ejemplo, si solicitas la ruta `/` de este sitio web, el controlador responsable de la página de inicio te producirá la página de inicio HTML tal y como la ves. Si estás construyendo una API REST, el controlador obtendrá o guardará algunos datos, y luego pedirá (amablemente) al motor de renderizado que produzca la respuesta apropiada.

En el caso de Buffalo, solemos llamar a los controladores "acciones".
## Definiendo una Acción

Las acciones (o controladores) de Buffalo son funciones de tipo {{< doclink href="github.com/gobuffalo/buffalo#Handler" message="Handler" >}}.

```go
func Home(c buffalo.Context) error {
	return c.Render(200, r.HTML("home.html"))
}
```

En este ejemplo, definimos una acción "Home", y le pedimos al motor de renderizado que produzca una página HTML usando la plantilla "home.html", y que responda con un código HTTP 200.

Cada acción toma un `buffalo.Context` como parámetro: Revisa [Contexto](/es/documentation/request_handling/context) para aprender todo lo que puedes hacer con él.

## Generando Acciones

Dado que escribir el código de las acciones es bastante redundante y repetitivo, Buffalo proporciona un generador para ayudarte.

```bash
$ buffalo g action --help
Generate new action(s)

Usage:
  buffalo generate action [name] [handler name...] [flags]

Aliases:
  action, a, actions

Flags:
  -d, --dry-run         dry run
  -h, --help            help for action
  -m, --method string   change the HTTP method for the generate action(s) (default "GET")
      --skip-template   skip generation of templates for action(s)
  -v, --verbose         verbosely run the generator
```

Para generar una acción para `users` sólo debes escribir:

```bash
$ buffalo g a users show index create
```

Esto generará los siguientes archivos:

```erb
├── actions/
│	├── users_test.go
│	└── users.go
│
└── templates/
	└── users/
		├── create.plush.html
		├── index.plush.html
		└── show.plush.html
```

Además, Buffalo registrará las rutas para `users` en el archivo `actions/app.go`:

```go
// actions/app.go

app.GET("/users/show", UsersShow)
app.GET("/users/index", UsersIndex)
app.GET("/users/create", UsersCreate)
```

En algunos casos necesitarás generar una acción con un método HTTP diferente a `GET`, para ese caso puedes usar el flag `--method`, como en el siguiente ejemplo:

```bash
$ buffalo g actions users message --method POST
```

En algunos otros escenarios necesitarás generar una acción sin generar una plantilla HTML (por ejemplo, para una API). Para omitir la generación de la plantilla HTML para crear una acción, puede pasar el flag --skip-template al generador, es decir:

```bash
$ buffalo g actions users update --skip-template
```

{{< note >}}
Este es el comportamiento por defecto para las aplicaciones generadas con el flag `--api`. Revisa [APIs](/documentation/guides/apis/) para mayor información.
{{< /note >}}

## Eliminando Acciones

Puede eliminar los archivos generados por este generador ejecutando:

```bash
$ buffalo destroy action users
```

O en forma abreviada:

```bash
$ buffalo d a users
```

## Siguientes Pasos

* [Recursos](/es/documentation/request_handling/resources) - Definir paquetes de acciones tipo CRUD.
* [Contexto](/es/documentation/request_handling/context) - Más información sobre Buffalo Context.
