<%= title("Additional Configuration") %>

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
