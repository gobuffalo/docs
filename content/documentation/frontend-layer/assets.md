---
Name: "Assets"
aliases:
  - /docs/assets
  - /en/docs/assets
---
# Assets

By default when a new Buffalo application is generated via the `buffalo new` command, a [Webpack](https://webpack.github.io) configuration file is generated, and the application is set up to use Webpack as the asset pipeline for the application.

If [`npm`](https://www.npmjs.com) is not found on the machine generating the new Buffalo application, then Webpack will not be configured and the asset pipeline would be skipped.

The asset pipeline can also be skipped during application generation with the `--skip-webpack` flag.

## JavaScript

The asset pipeline is initially configured to support ES6 JavaScript files, with `/assets/js/application.js` being the main entry point.

The following are automatically installed and configured during setup of the asset pipeline:

* [jQuery](https://jquery.com)
* [Bootstrap 4](http://getbootstrap.com)
* [jQuery UJS](https://github.com/rails/jquery-ujs)

None of the installed packages are required, and may be removed. They are included for convenience.

## CSS

By default the asset pipeline is configured to use [.scss](http://sass-lang.com) files, with `/assets/css/application.scss` as the main entry point. This, of course, can be changed.

The following are automatically installed and configured during setup of the asset pipeline:

* [Bootstrap 4](http://getbootstrap.com)
* [Font Awesome](http://fontawesome.io)

None of the installed packages are required, and may be removed. They are included for convenience.

## Other Assets

Any assets placed in the `/assets` folder will be copied to the "distribution" automatically, and can be found at `/assets/path/to/asset`.

## Asset Fingerprinting

In `v0.9.5` asset fingerprinting was introduced to the default Webpack configuration for new applications. Asset fingerprinting works by generating a hash of the file contents and appending it to the name of the file. So, for example, `application.js` might be come `application.a8adff90f4c6d47529c4.js`. The benefit of this is that it allows for assets to be cached but still allow for that cache to be broken when a change has been made to the contents of this file.

Note that in order for this to work, buffalo will expect a `/public/assets/manifest.json` file to be present, containing the mappings between the files you reference in the helpers (eg `javascriptTag("application.js")`) and their hashed counterparts. This is not something you need to worry about if you are using the default Webpack configuration. However, if you choose to use the `--skip-webpack` flag when generating the project, keep in mind you will have to handle that yourself. While having the manifest file present is not strictly required for your application to run, you may experience caching problems without it during development.

**NOTE:** Applications written before `v0.9.5` may need to set an `AssetsBox` in their `actions/render.go` file in the `render.Options`, if assets are not rendering properly. It is recommended to move the one from the `actions/app.go` file into there instead. This will not setup asset finger printing, but will make sure the assets render correctly. See changes [https://github.com/gobuffalo/docs/commit/00673ab3446a9a7209bbd243e4594bd679c81a69#diff-c1ebdbddf205da1687721a8acd29043cR43](https://github.com/gobuffalo/docs/commit/00673ab3446a9a7209bbd243e4594bd679c81a69#diff-c1ebdbddf205da1687721a8acd29043cR43) and [https://github.com/gobuffalo/docs/commit/00673ab3446a9a7209bbd243e4594bd679c81a69#diff-25015af78e14806bd828e39a29a403fbR13](https://github.com/gobuffalo/docs/commit/00673ab3446a9a7209bbd243e4594bd679c81a69#diff-25015af78e14806bd828e39a29a403fbR13) for examples.

By default new applications are setup to fingerprint only JavaScript and CSS files.

## Asset Helpers

With the introduction of asset fingerprinting in `v0.9.5` it became difficult to find asset files because the name of the file kept changing. To help with this, three new helpers were introduced.

1. `assetPath` - This helper will return the path of the requested asset. For example, `<%= assetPath("application.js") %>` would return something like `/assets/application.a8adff90f4c6d47529c4.js`.

2. `javascriptTag` - This helper will generate a `<script src="xxx"></script>` style tag for the requested JavaScript file. Example: `<%= javascriptTag("application.js") %>` would return something like `<script src="/assets/application.bd76587ded82386f388f.js" type="text/javascript"></script>`.

3. `stylesheetTag` - This helper will generate a `<link href="xxx">` style tag for the requested CSS file. Example: `<%= stylesheetTag("application.css") %>` would return something like `<link href="/assets/application.bd76587ded82386f388f.css" media="screen" rel="stylesheet" />`.

## Building Assets in Development

The `buffalo dev` command, in addition to watching and rebuilding the application's Go binary, will watch, and rebuild the asset pipeline as well. Nothing special needs to be run.

## Building Assets for Deployment

The `buffalo build` command will build the asset pipeline, and properly attach it to the generated Go binary. One binary to run them all! See [Packing](/documentation/deploy/packing) for more options on building assets for deployment.
