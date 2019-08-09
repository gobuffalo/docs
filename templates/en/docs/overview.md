# Overview

Welcome aboard! While Buffalo can be considered as a framework, it's mostly an ecosystem of Go and Javascript libraries curated to fit together. Most of these components can be switched for another, but we'll only provide support for this default mix.

In this chapter, we'll make a tour of the default bricks shipped with your Buffalo app.

## Backend libraries

### buffalo

Buffalo is the "glue" between all the provided components. It wraps the libraries and manages the workflow.

### gorilla/mux

[gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) is one of the most used routers in Go. While some routers are faster (like [httprouter](https://github.com/julienschmidt/httprouter)), gorilla/mux is the one providing the most features while being fast enough.

### pop

[pop](https://github.com/gobuffalo/pop) is the default ORM for Buffalo. It provides the `soda` toolbox to help you with your database needs and supports several databases, such as PostgreSQL, MySQL and SQLite.

### plush

[plush](https://github.com/gobuffalo/plush) is the default templating engine for Buffalo. Its syntax is close to ERB templates (in Ruby).

### packr

[packr](https://github.com/gobuffalo/packr) is a Go bundler for your static assets (templates, images and so on). Packr aims to produce a final binary with everything embedded in it.

## Frontend libraries

### Bootstrap

[Bootstrap](https://getbootstrap.com/) is one of the most famous frontend toolkit library. It helps to build responsive interfaces using common components like tables, carousels or grid layouts.

### jQuery

[jQuery](https://jquery.com/) is a rich library aiming to make DOM manipulation and AJAX queries simple. While it's less used now, many projects still have it as a side-companion to help supporting all the browsers.

### Webpack

[Webpack](https://webpack.js.org/) is a well-known Javascript assets bundler. It will take care of your Javascript, CSS, images and static assets files.

Webpack is configured by default to hash and minify your assets.

## Next Steps

* [Installation](/en/docs/getting-started/installation) - Install Buffalo!