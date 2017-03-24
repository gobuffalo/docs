<%= title("Running Your Application in Development", {name: "running-in-dev"}) %>

_Before starting Buffalo for the first time, please head over to the [Database](/docs/db) docs and read a little bit about setting up your databases._

One of the downsides to Go development is the lack of code "reloading". This means as you change your code you need to manually stop your application, rebuild it, and then restart it. Buffalo finds this is annoying, and wants to make life better for you.

<%= code("text") { %>
$ buffalo dev
<% } %>

The `dev` command will watch your `.go` and `.html` files and the [asset](/docs/assets) folder by default. It will rebuild and restart your binary for you automatically so you don't have to worry about such things.

Just run the `buffalo dev` command and go to [localhost:3000/](http://localhost:3000/) to see all changes live!
