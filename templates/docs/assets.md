# Assets

By default when a new Buffalo application is generated via the `buffalo new` command, a [Webpack](https://webpack.github.io) configuration file is generated, and the application is set up to use Webpack as the asset pipeline for the application.

If [`npm`](https://www.npmjs.com) is not found on the machine generating the new Buffalo application, then Webpack will not be configured and the asset pipeline would be skipped.

The asset pipeline can also be skipped during application generation with the `--skip-webpack` flag.

{{ partial "topics.html" }}

{{#panel title="JavaScript"}}

The asset pipeline is initially configured to support ES6 JavaScript files, with `/assets/js/application.js` being the main entry point.

The following are automatically installed and configured during setup of the asset pipeline:

* [jQuery](https://jquery.com)
* [Bootstrap](http://getbootstrap.com)
* [jQuery UJS](https://github.com/rails/jquery-ujs)

None of the installed packages are required, and may be removed. They are included for convenience.
{{/panel}}

{{#panel title="CSS"}}

By default the asset pipeline is configured to use [.scss](http://sass-lang.com) files, with `/assets/css/application.scss` as the main entry point. This, of course, can be changed.

The following are automatically installed and configured during setup of the asset pipeline:

* [Bootstrap](http://getbootstrap.com)
* [Font Awesome](http://fontawesome.io)

None of the installed packages are required, and may be removed. They are included for convenience.

{{/panel}}

{{#panel title="Other Assets"}}

Any assets placed in the `/assets` folder will be copied to the "distribution" automatically, and can be found at `/assets/path/to/asset`.

{{/panel}}

{{#panel title="Building Assets in Development"}}

The `buffalo dev` command, in addition to watching and rebuilding the application's Go binary, will watch, and rebuild the asset pipeline as well. Nothing special needs to be run.

{{/panel}}

{{#panel title="Building Assets for Deployment"}}

The `buffalo build` command will build the asset pipeline, and properly attach it to the generated Go binary. One binary to run them all!

{{/panel}}
