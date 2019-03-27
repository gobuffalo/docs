<% seoDescription("How to define and use Action Controllers?") %>
<% seoKeywords(["buffalo", "go", "golang", "actions", "controller", "generator"]) %>

<% contentFor("docsLanding") { %>
  <%= partial("docs/landing.html", {title: "Actions"}) %>
<% } %>

In this chapter, you'll learn how action controllers work; and how you can generate them using the built-in generators.

## What is a Controller?

Controllers are the *C* part of the [MVC pattern](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller). They handle the logic given the router decision, and produce an appropriate response.

For instance, if you request the `/` path of this website, the handler responsible of the home page will produce you the HTML home page as you see it. If you're building a REST API, the controller will fetch or save some data, then ask (politely) the render engine to produce the appropriate response.

In Buffalo case, we commonly call controllers "actions".

## Define an Action

Buffalo's actions (or controllers) are <%= doclink("github.com/gobuffalo/buffalo#Handler") %> functions:

```go
func Home(c buffalo.Context) error {
	return c.Render(200, r.HTML("home.html"))
}
```

In this example, we defined a "Home" action, and asked the rendering engine to produce an HTML page using the "home.html" template, and to reply with an HTTP 200 code.

Each action takes a `buffalo.Context` as parameter: see [Context](/en/docs/context) to learn more about all you can do with it.

<%= partial("docs/generators/action.md") %>

## Next Steps

* [Resources](/en/docs/resources) - Define CRUD-like action bundles.
* [Context](/en/docs/context) - Learn more about Buffalo Context.