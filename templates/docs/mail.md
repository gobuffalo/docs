<% contentFor("docsLanding") { %>
  <%= partial("docs/landing.html", {title: "Mailers"}) %>
<% } %>

E-mails are part of most (web) systems. It's a standard way to inform users about platform news, confirmations, marketing stuff, and so on.

Buffalo provides, out of the box, a mailer extension with a standard SMTP sender. A generator is included, to allow you to work with emails as fast as possible.

<%= partial("docs/mail/generator.md") %>
<%= partial("docs/mail/example.md") %>
<%= partial("docs/mail/context.md") %>
<%= partial("docs/mail/configuration.md") %>

## Sender Implementations

Some alternate [`Sender`](https://godoc.org/github.com/gobuffalo/buffalo/mail#Sender) implementations are provided by the Buffalo community:

* [postmark-sender](https://github.com/paganotoni/postmark-sender) - A sender to work with [Postmark](https://postmarkapp.com/). Author: [@paganotoni](https://github.com/paganotoni).
* [mocksmtp](https://github.com/stanislas-m/mocksmtp) - A mock implementation that can be used for testing. Author: [@stanislas-m](https://github.com/stanislas-m).
* [sendgrid-sender](https://github.com/paganotoni/sendgrid-sender) - A sender to work with [Sendgrid](https://sendgrid.com/). Author: [@paganotoni](https://github.com/paganotoni).
* [mailopen](https://github.com/paganotoni/mailopen) - A sender that opens emails in browser. Author: [@paganotoni](https://github.com/paganotoni).
