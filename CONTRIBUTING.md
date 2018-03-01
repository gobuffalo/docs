Contribute to Buffalo documentation
===================================

This documentation website is a team effort! Making sure people can understand how to use Buffalo is the aim of this project.
Feel free to join us, even to fix a small typo.

## Translations guidelines
The current available languages are:
* English (upstream version)
* French

### Adding a new language
Before translating the docs into another language, make sure the french version of the page is available. If it's not, this probably means the page will move a bit: it's better to wait for a stable version, so you don't need to translate the page for (almost) nothing.

### Internationalization
Internationalization (or i18n) is the way to mark software parts as translatable. Buffalo docs i18n uses two ways to do so:
* The `t` helper is used to manage translations for fixed contents: menu titles, for instance.
* The Localized views are used to handle most of the documentation pages: this allows people to provide customized versions for the different supported languages, and it's a more efficient way to proceed, since the whole page needs to be translated.

See https://gobuffalo.io/docs/localization for further info about this two tools.

## Hack the code

This website is powered by [Buffalo](https://github.com/gobuffalo/buffalo), so if your read this docs, you'll be able to run it and make the changes you want! (if you can't, the docs are not good enough, and are waiting for a [PR](https://github.com/gobuffalo/gobuffalo/pulls)! ;)

### Requirements

* `npm` & `yarn` for the frontend pipeline.
* A working go (1.8.1 or greater) setup.

### Run the website for dev
```bash
$ go get github.com/gobuffalo/gobuffalo
$ cd $GOPATH/github.com/gobuffalo/gobuffalo
$ buffalo dev
```

And that's it!