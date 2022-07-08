---
name: Cookies
seoDescription: "Cookies"
seoKeywords: ["buffalo", "go", "golang", "http", "cookie"]
weight: 9
aliases:
  - /docs/cookies
  - /es/docs/cookies
---

# Cookies

Una cookie HTTP es un pequeño dato que un servidor envía al navegador web del usuario. El navegador puede almacenar estos datos y enviarlos de nuevo al mismo servidor, incluso después de reiniciar el navegador (a diferencia de una [sesión de navegador](/es/documentation/request_handling/sessions)).

Las cookies (HTTP) se utilizan comúnmente para guardar el estado de los usuarios (como por ejemplo, si el usuario se ha conectado). Ver [https://golang.org/pkg/net/http/#Cookie](https://golang.org/pkg/net/http/#Cookie) para más información sobre las cookies en Go.

## Setting a Cookie

```go
func MyHandler(c buffalo.Context) error {
  // ...
  c.Cookies().Set("user_id", user.ID, 30 * 24 * time.Hour)
  // ...
}
```

## Cómo definir una cookie

```go
func MyHandler(c buffalo.Context) error {
  // ...
  exp := time.Now().Add(365 * 24 * time.Hour) // expire in 1 year
  c.Cookies().SetWithExpirationTime("user_id", user.ID, exp)
  // ...
}
```

## Cómo definir una cookie con Ruta

```go
func MyHandler(c buffalo.Context) error {
  // ...
  c.Cookies().SetWithPath("user_id", user.ID, "/user")
  // ...
}
```

## Manera Avanzada de Definir una Cookie

```go
import "net/http"
```

```go
func MyHandler(c buffalo.Context) error {
  // ...
  ck := http.Cookie{
    Name:    "token",
    Value:   token,
    Path:    "/",
    Expires: time.Now().Add(30 * 24 * time.Hour), // expire in 1 month
  }

  http.SetCookie(c.Response(), &ck)
  // ...
}
```

Ver [Cookie struct](https://golang.org/src/net/http/cookie.go) para conocer otros parámetros.

## Obtener una Cookie

```go
func MyHandler(c buffalo.Context) error {
  value, err := c.Cookies().Get("user_id")
  if err != nil {
    return err
  }
  return c.Render(200, r.String(value))
}
```

## Eliminando una Cookie


```go
func MyHandler(c buffalo.Context) error {
  c.Cookies().Delete("user_id")
  // ...
}
```
