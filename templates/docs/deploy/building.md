<% seoDescription("Packing") %>
<% seoKeywords(["buffalo", "go", "golang", "pack", "binary", "build"]) %>

<%= h1("Packing") %>

Now, your project is ready to be deployed. In this section, you will learn how to package a version of your app to deploy it on a server.

<%= title("The build command") %>

Buffalo features a command, `build`, that will build a **full binary** of your application including, but not limited to; assets, migrations, templates, etc. If you buy into the “Buffalo Way”, things just work. It's a wonderful experience. :)

<%= partial("docs/building/build_trace.md") %>

When the build finishes, you have a fresh baked binary in the `bin` folder. It will also have the **compilation time** and the **git commit SHA** burnt in, thus making the binaries “versioned”.

<%= title("Customize the build") %>

To get the list of available options, use the help command:

<%= partial("docs/building/build_options.md") %>

### Binary name / location

By default, your application will be built in the `bin` directory of your project, and the name of the executable will be the name you used to create the project with the `new` command.

You can change this default name by using the `-o` or `-output` flag:

<%= partial("docs/building/output_flag.md") %>

In fact, you can change the target directory too:

<%= partial("docs/building/output_dir.md") %>

### Extract Assets in a Zip File

By default, your whole app is packed into a single executable, assets included. In production setups, you may want to serve these assets with a proxy server (like Apache or NGINX), to lower the app load. You may even use a *CDN* to handle your assets.

Buffalo provides a way to extract compiled app assets into a single archive, using the `-e` or `-extract-assets` flag:

<%= partial("docs/building/extract_assets.md") %>

By default, the assets archive is put in the *bin* directory, but if you change the executable output directory with the `-o` flag, the assets will be put in the same directory.

<%= partial("docs/building/extract_assets_layout.md") %>

<%= title("Advanced Options") %>

### Building “Static”/CGO Binaries

Building statically linked binaries that contain CGO, think SQLite3, can be tricky. By using the `--static` flag with `buffalo build`, the flags `--ldflags '-linkmode external -extldflags "-static"'` will be added to the `go build` command.

### Build Tags

When building a Buffalo binary using the `buffalo build` command, you can pass `--tags` and `--ldflags` to the built binary; just as you normally would when using the `go build` tools.

```bash
$ buffalo build --tags="mytag" --ldflags="-X foo.Bar=baz"
```

<%= title("Binary Commands") %>

### Modes
Binaries, by default, run in `development` mode, which means all of the sub-commands will run in that mode as well. To change the mode, you must use the `GO_ENV` environment variable.

```bash
$ GO_ENV=production ./coke
```

### Available commands

Once a binary has been built, there are several sub-commands that can be run on that binary:

#### Default

The default command, if you just run the binary, will start the application.

#### migrate

The `migrate` sub-command will run the migrations for the application.

#### version

The `version` sub-command will output the version information for the binary, including the name, the git commit SHA used to build the binary, and the time the binary was built.

```bash
$ ./coke version
coke version 69b6a8b ("2017-04-03T10:19:46-04:00")
```

#### task

The `task` sub-command runs tasks.

```bash
$ ./coke task greet

Hello World!
```

<%= title("Next Steps") %>

* [Using a Proxy](/en/docs/deploy/proxy) - Integrate your app with a server like NGINX.
* [Systemd Service](/en/docs/deploy/systemd) - Run your app as a systemd service.
* [Cloud Providers](/en/docs/deploy/providers) - Deploy your app on a cloud provider.