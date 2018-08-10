<%= title("Running Your Application in Development") %>

<%= note() { %>
Before starting Buffalo for the first time, please head over to the [Database](/docs/db) docs and read a little bit about setting up your databases.
<% } %>

One of the downsides to Go development is the lack of code “reloading”. This means as you change your code **you need to manually stop** your application, rebuild it, then restart it. Buffalo finds this is annoying, and wants to make life better for you.

```bash
$ buffalo dev
```

The `dev` command will watch your `.go` and `.html` files and the [asset](/docs/assets) folder by default. It will **rebuild and restart your binary for you** automatically, so you don't have to worry about such things.

Just run the `buffalo dev` command and go to [localhost:3000/](http://localhost:3000/) to see all changes live!

<figure>
  <img src="/assets/images/new-coke.png" title="screenshot">
  <figcaption>The brand new Coke app.</figcaption>
</figure>

#### Run the dev server on a custom port

Sometimes you will already have an app working on the 3000 port. You can configure the dev server port by providing the `PORT` environment variable:

```bash
$ PORT=3001 buffalo dev
```

You can also take a look at the [Env Variables](/docs/env-vars) chapter for further information on Buffalo configuration.