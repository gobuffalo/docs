<%= h1("Localization") %>

Translating your app is an effective way to **make it understandable to the many people** around the globe! Buffalo uses the [go-i18n](github.com/nicksnyder/go-i18n) project to provide the <abbr title="internationalization">i18n</abbr> (providing default text, often in english) and <abbr title="localization">l10n</abbr> (providing translation for the other supported languages) of your app.

<%= title("Markup the translatable strings") %>

<%= note() { %>
<%= partial("docs/disclaimer.html") %>
<% } %>

Translatable strings must use a **specific markup** to allow the l10n engine to get the translations.

In a plush template, you can use the `t` (singular strings) and <%= sinceVersion("0.10.0") %> the `tp` (singular & plural) helpers:

```html
&lt;%= t("greetings") %&gt;
```

<%= sinceVersion("0.10.0") %>

```html
&lt;%= tp("messages-notification", notificationsLen) %&gt;
```

<%= title("Provide translations") %>

Translations are stored in the `locales` folder. By default, they are stored in a `all.en-us.yaml` file for the american english strings.

You can provide translations for another language by providing a new file `all.my-language-code.yaml`. If you want to split your strings into logical modules, you can even create multiples files, e.g. `users.en-us.yaml` for the user-related stuff, and `all.en-us.yaml` for the global stuff.

The localization format used by [go-i18n](github.com/nicksnyder/go-i18n) is the following:

```yaml
- id: greetings
  translation: "Welcome to Buffalo (EN)"

- id: messages-notification
  translation:
    one: "You have {{.Count}} notification"
    other: "You have {{.Count}} notifications"
```

<%= title("Define the default language") %>

To define the default language of your app, you need to edit the `app.go` file in the `actions` folder:

```go
// Setup and use translations:
var err error
if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
  app.Stop(err)
}
app.Use(T.Middleware())
```

Changing `"en-US"` to another language code will change the default language.


<%= title("Translate strings in actions") %>

You can use `T` from within actions:

```go
func MyActionHandler(c buffalo.Context) error {
  msg, err := T.Translate(c, "greetings")
  if err != nil {
	  return err
	}
	c.Flash().Add("info", msg)
	return c.Render(200, r.HTML("action.html"))
}
```
