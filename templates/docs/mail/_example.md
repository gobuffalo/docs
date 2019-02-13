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
