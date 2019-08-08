## Custom Asset Pipeline Commands

This release replaces the hard coded Webpack commands (`dev` & `build`) with standard `package.json` scripts which can be called with either NPM or Yarn.

The new applications will be generated with the following new section in the `package.json`:

```json
"scripts": {
    "build": "webpack -p --progress",
    "dev": "webpack --watch"
}
```

Buffalo will then call `yarn run` build (or `npm run build`) when you call `buffalo build`; and it will call `yarn run dev` (or `npm run dev`) when you call `buffalo dev`. This allows you to customize the scripts Buffalo calls on these steps, providing custom arguments to the Webpack command or even running your own tool chain instead.

If the `build` or `dev` scripts can't be found, Buffalo will fall back on the old behavior running the hard coded Webpack commands.
