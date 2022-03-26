---
name: Mailers
---

E-mails are part of most (web) systems. It's a standard way to inform users about platform news, confirmations, marketing stuff, and so on.

Buffalo provides, out of the box, a mailer extension with a standard SMTP sender. A generator is included, to allow you to work with emails as fast as possible.

## Generator

When the generator is run for the first time it will bootstrap a new `mailers` package and a new `templates/mail` directory.

```bash
$ buffalo generate mailer welcome_email
```


## Example Usage

```go
// mailers/mail.go
package x

import (
  "log"

  "github.com/gobuffalo/buffalo/render"
  "github.com/gobuffalo/envy"
  "github.com/gobuffalo/packr"
  "github.com/gobuffalo/plush"
  "github.com/gobuffalo/buffalo/mail"
  "github.com/pkg/errors"
  "gitlab.com/wawandco/app/models"
)

var smtp mail.Sender
var r *render.Engine

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

  // The rendering engine, this is usually generated inside actions/render.go in your buffalo app.
  r = render.New(render.Options{
    TemplatesBox:   packr.NewBox("../templates"),
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

This `SendContactMessage` could be called by one of your actions, i.e. the action that handles your contact form submission.

```go
// actions/contact.go
...

func ContactFormHandler(c buffalo.Context) error {
  contact := &models.Contact{}
  c.Bind(contact)

  // Calling to send the message
  SendContactMessage(contact)
  return c.Redirect(302, "contact/thanks")
}
...
```

You can add your own custom plush functions by binding them in as data.

```go
func UUIDToString(u uuid.UUID) string {
  return fmt.Sprintf("%s", u)
}

  m := mail.NewMessage()
  ...
  
  // Data that will be used inside the templates when rendering.
  data := map[string]interface{}{
    "contact": c,
    "UUIDToString": UUIDToStringHelper,
  }
```


## Using Context Variables

<%= sinceVersion("0.13.0-rc1") %>

To use context variables such as [RouteHelpers](/en/docs/routing#using-route-helpers-in-templates) or those set with
`c.Set(...)`, `mail.New` accepts a `buffalo.Context`.

```go
func SendMail(c buffalo.Context) error {
  m := mail.New(c)
  ...

  m.AddBody(r.HTML("mail.html"))
  return SMTP.Send(m)
}
```

```html
&lt;a href="\<%= awesomePath() %>">Click here&lt;/a>
```


## Additional Configuration

If you're using Gmail or need to configure your SMTP connection, you can use the `Dialer` property on the SMTPSender, p.e: (for Gmail)

```go
// mailers/mail.go
...
var smtp mail.Sender

func init() {
  port := envy.Get("SMTP_PORT", "465")
  // or 587 with TLS

  host := envy.Get("SMTP_HOST", "smtp.gmail.com")
  user := envy.Get("SMTP_USER", "your@email.com")
  password := envy.Get("SMTP_PASSWORD", "yourp4ssw0rd")

  // Assigning to smtp later to preserve type
  var err error
  sender, err := mail.NewSMTPSender(host, port, user, password)
  sender.Dialer.SSL = true

  //or if TLS
  sender.Dialer.TLSConfig = &tls.Config{...}

  smtp = sender
}
...
```


## Sender Implementations

Some alternate [`Sender`](https://godoc.org/github.com/gobuffalo/buffalo/mail#Sender) implementations are provided by the Buffalo community:

* [postmark-sender](https://github.com/paganotoni/postmark-sender) - A sender to work with [Postmark](https://postmarkapp.com/). Author: [@paganotoni](https://github.com/paganotoni).
* [mocksmtp](https://github.com/stanislas-m/mocksmtp) - A mock implementation that can be used for testing. Author: [@stanislas-m](https://github.com/stanislas-m).
* [sendgrid-sender](https://github.com/paganotoni/sendgrid-sender) - A sender to work with [Sendgrid](https://sendgrid.com/). Author: [@paganotoni](https://github.com/paganotoni).
* [mailopen](https://github.com/paganotoni/mailopen) - A sender that opens emails in browser. Author: [@paganotoni](https://github.com/paganotoni).
