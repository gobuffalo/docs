{{#panel title="Running Your Application in Development" name="running-in-dev"}}

_Before starting Buffalo for the first time, please head over to the [Database](/docs/db) docs and read a little bit about setting up your databases._

One of the downsides to Go development is the lack of code "reloading". This means as you change your code you need to manually stop your application, rebuild it, and then restart it. Buffalo finds this is annoying, and wants to make life better for you.

```
$ buffalo dev
```

The `dev` command will watch your `.go` and `.html` files by default. It will rebuild and restart your binary for you so you don't have to worry about such things. The `dev` command will also watch [asset](/docs/assets) and recompile those as well.

Just run the `buffalo dev` command and start coding.

{{/panel}}
