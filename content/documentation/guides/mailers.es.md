---
name: Mailers
weight: 13
aliases:
  - /documentation/guides/mail
  - /docs/mail
  - /es/docs/mail
---

# Mailers

Los correos electrónicos son parte de la mayoría de los sistemas (web). Son una forma estandar de informar a los usuarios sobre las noticias de la plataforma, conformaciones, cosas de marketing, etc.

Buffalo proporciona, lista para usar, una extension de correos con un remitenter SMTP. Se incluye un generador, para permitirte trabajar con emails lo más rapido posible.


## Generador

Cuando el generador se ejecuta por primera vez, se creará un nuevo paquete `mailers` y un nuevo directorio `templates/mail`.

```bash
$ buffalo generate mailer welcome_email
```


## Ejemplo de uso

```go
// mailers/mailers.go

package mailers

import (
	"log"

	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr/v2"
)

var (
	smtp mail.Sender
	r    *render.Engine
)

func init() {
	// Pulling config from the env.
	port := envy.Get("SMTP_PORT", "1025")
	host := envy.Get("SMTP_HOST", "localhost")
	user := envy.Get("SMTP_USER", "")
	password := envy.Get("SMTP_PASSWORD", "")

	var err error
	smtp, err = mail.NewSMTPSender(host, port, user, password)

	if err != nil {
		log.Fatal(err)
	}

	r = render.New(render.Options{
		HTMLLayout:   "layout.html",
		TemplatesBox: packr.New("app:mailers:templates", "../templates/mail"),
		Helpers:      render.Helpers{},
	})
}


// SendContactMessage Sends contact message to contact@myapp.com
func SendContactMessage(c *models.Contact) error {

  // Creates a new message
  m := mail.NewMessage()
  m.From = "sender@myapp.com"
  m.Subject = "New Contact"
  m.To = []string{"contact@myapp.com"}

  // Data that will be used inside the templates when rendering.
  data := map[string]interface{}{
    "contact": c,
  }

  // You can add multiple bodies to the message you're creating to have content-types alternatives.
  err := m.AddBodies(data, r.HTML("mail/contact.html"), r.Plain("mail/contact.txt"))

  if err != nil {
    return errors.WithStack(err)
  }

  err = smtp.Send(m)
  if err != nil {
    return errors.WithStack(err)
  }

  return nil
}
```
Este `SendContactMessage` podría ser llamado por una de tus acciones, es decir, la acción que maneja el envío de tu formulario de contacto.

```go
// actions/contact.go

func ContactFormHandler(c buffalo.Context) error {
  contact := &models.Contact{}
  if err := c.Bind(contact); err != nil {
    // ...
  }
  // ...

  // Calling to send the message
  SendContactMessage(contact)
  return c.Redirect(http.StatusFound, "contact/thanks")
}
```

Puedes agregar tus propias funciones de plush personalizadas vinculándolas como datos.

```go
// UUIDToString helper func
func UUIDToString(u uuid.UUID) string {
  return fmt.Sprintf("%s", u)
}
```


``` go
// SendContactMessage Sends contact message to contact@myapp.com
func SendContactMessage(c *models.Contact) error {

  // Creates a new message
  m := mail.NewMessage()

  // ...

  data := map[string]interface{}{
    "contact": c,
    "UUIDToString": UUIDToStringHelper,
  }

  // ...
}
```

## Usando variables de Contexto

{{< since "0.13.0-rc1" >}}

Para usar variables de contexto como [RouteHelpers](/documentation/request_handling/routing#using-route-helpers-in-templates) o los establecidos con `c.Set(...)`; `mail.New` acepta `buffalo.Context`.

```go
func SendMail(c buffalo.Context) error {
  m := mail.New(c)
  // ...

  m.AddBody(r.HTML("mail.html"))
  return SMTP.Send(m)
}
```


## Configuración adicional

Si estas usando Gmail o necesitas configurar tu conexión SMTP, puedes usar la propiedad `Dialer` en el SMTPSender, por ejemplo (para Gmail)

```go
// mailers/mailers.go

package mailers

// ...

var (
	smtp mail.Sender
	r    *render.Engine
)

func init() {
	port := envy.Get("SMTP_PORT", "465") // or 587 with TLS
	host := envy.Get("SMTP_HOST", "smtp.gmail.com")
    user := envy.Get("SMTP_USER", "your@email.com")
    password := envy.Get("SMTP_PASSWORD", "yourp4ssw0rd")

	var err error
    // sender se asignará a smtp despues de configurarse
	sender, err = mail.NewSMTPSender(host, port, user, password)
    sender.Dialer.SSL = true

    // o si es TLS
    sender.Dialer.TLSConfig = &tls.Config{...}

	if err != nil {
		log.Fatal(err)
	}

    smtp = sender

	// ...
}
```


## Implementaciones de Sender

Algunas alternativas de implementaciones de [`Sender`](https://godoc.org/github.com/gobuffalo/buffalo/mail#Sender) proporcionadas por la comunidad de Buffalo:

| Librería                                                         | Descripción                                                    | Autor                                          |
|:-----------------------------------------------------------------|:---------------------------------------------------------------|:-----------------------------------------------|
| [postmark-sender](https://github.com/paganotoni/postmark-sender) | Un sender que trabaja con [Postmark](https://postmarkapp.com/) | [@paganotoni](https://github.com/paganotoni)   |
| [mocksmtp](https://github.com/stanislas-m/mocksmtp)              | Una implementación simulada que se puede usar para pruebas     | [@stanislas-m](https://github.com/stanislas-m) |
| [sendgrid-sender](https://github.com/paganotoni/sendgrid-sender) | Un sender que trabaja con [Sendgrid](https://sendgrid.com/)    | [@paganotoni](https://github.com/paganotoni)   |
| [mailopen](https://github.com/paganotoni/mailopen)               | Un sender que aber los correos en el navegador                 | [@paganotoni](https://github.com/paganotoni)   |

