Contribute to Buffalo documentation
===================================

This documentation website is a team effort! Making sure people can understand how to use Buffalo is the aim of this project.
Feel free to join us, even to fix a small typo.

## Writing style guidelines
* Keep the tone as neutral as possible. The reader is not a friend, but deserves respect though: write facts, not feelings.
* Prefer the formal form when you address to the reader: in french, it means using "vous" instead of "tu" ; in german "Sie" instead of "Du".
* Don't assume the reader is a male, prefer a more inclusive form. In french, though, **we don't use the inclusive writing form** because it's too political right now and there is no consensus on the subject.
* Don't drop code without any explanation.
* We don't assume the reader masters Go. Explain any code you provide but make it simple.
* If you're documenting a new feature, provide at least the english version. You're not required to provide the translations for the others versions, but you can. :)
* Document only Buffalo-specific things. Put a link to the proper docs for external things (e.g. Buffalo provides generators using `nodejs` or `yarn`: don't document how to use these tools, refer to the proper docs).
* Docs are split into several parts (see the nav menu). If you're not sure where to put a new doc, ask @gobuffalo/docs-managers.
* When you provide an example, try to use the same theme in the same page. If you gave a music-related example in the previous section, try to stay in the music theme for the remaining of the chapter.

## Translations guidelines
The current available languages are:
* English (upstream version)
* French

### Adding a new language
Before translating the docs into another language, **make sure the French version of the page is available**. If it's not, this probably means the page will move a bit: it's better to wait for a stable version, so you don't need to translate the page for (almost) nothing.

### Internationalization
Internationalization (or i18n) is the way to mark software parts as translatable. Buffalo docs i18n uses two ways to do so:
* The `t` helper is used to manage translations for fixed contents: menu titles, for instance.
* The Localized views are used to handle most of the documentation pages: this allows people to provide customized versions for the different supported languages, and it's a more efficient way to proceed, since the whole page needs to be translated.

See https://gobuffalo.io/docs/localization for further info about this two tools.

## Hack the code

The site is powered by [Buffalo](https://github.com/gobuffalo/buffalo) itself, so if you're reading the docs, you'll be able to run them and make the changes you want! (If you can't, the docs are not good enough, and are waiting for a [PR](https://github.com/gobuffalo/docs/pulls)! ;))

### Requirements

* Node.js >= 8.x
* `npm` or `yarn` for the frontend pipeline.
* [Hugo](https://gohugo.io/getting-started/installing/)

### Run the website for dev
```console
$ git clone git@github.com:gobuffalo/docs.git
$ cd docs
$ npm install
$ hugo server -D
```

And that's it!
