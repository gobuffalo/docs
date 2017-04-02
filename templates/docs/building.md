# Building

As explained in the [Getting started](/docs/getting-started) section, you can use the `build` command to build a full binary of your application:

<%= code("text") { %>
$ buffalo build
<% } %>

<%= code("text") { %>
Buffalo version 0.8.0

--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o bin/coke -ldflags -X main.version=b5dffda -X main.buildTime="2017-03-20T11:05:23-04:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go

<% } %>

<%= title("Define a custom binary name", {name: "custom-bin-name"}) %>

By default, your application will be built in the `bin` directory of your project, and the name of the executable will be the name you used to create the project with the `new` command.

You can change this default name by using the `-o` or `-output` flag:

<%= code("text") { %>
$ buffalo build -o bin/cookies
<% } %>

<%= code("text") { %>
Buffalo version 0.8.0

--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o bin/cookies -ldflags -X main.version="2017-04-02T08:32:28+02:00" -X main.buildTime="2017-04-02T08:32:28+02:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go

<% } %>

In fact, you can change the target directory too:

<%= code("text") { %>
$ # Put the app in my home directory, as "coke"
$ buffalo build -o ~/coke
<% } %>

<%= code("text") { %>
Buffalo version 0.8.0

--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o ~/coke -ldflags -X main.version="2017-04-02T08:32:28+02:00" -X main.buildTime="2017-04-02T08:32:28+02:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go

<% } %>

<%= title("Extract assets in a zip file", {name: "extract-assets"}) %>

By default, your whole app is packed into a single executable, assets included. In production setups, you may want to serve these assets with a proxy server (like Apache or NGINX), to lower the app load. You may even use a *CDN* to handle your assets.

Buffalo provides a way to extract compiled app assets into a single archive, using the `-e` or `-extract-assets` flag:

 <%= code("text") { %>
$ buffalo build -e
<% } %>

<%= code("text") { %>
Buffalo version 0.8.0

--> cleaning up target dir
--> running node_modules/.bin/webpack
--> build assets archive
--> disable self assets handling
--> running go build -v -o bin/coke -ldflags -X main.version="2017-04-02T08:45:58+02:00" -X main.buildTime="2017-04-02T08:45:58+02:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go

<% } %>

Please note this will disable the internal assets handling too, so the final executable is lighter.

By default, the assets archive is put in the *bin* directory, but if you change the executable output directory with the `-o` flag, the assets will be put in the same directory.