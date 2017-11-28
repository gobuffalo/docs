<%= h1("Localization") %>

Translating your app is an effective way to **make it understandable to the many people** around the globe! Buffalo uses the [go-i18n](https://github.com/nicksnyder/go-i18n) project to provide the <abbr title="internationalization">i18n</abbr> (providing default text, often in english) and <abbr title="localization">l10n</abbr> (providing translation for the other supported languages) of your app.

<%= title("Markup the translatable strings") %>

<%= note() { %>
<%= partial("docs/disclaimer.html") %>
<% } %>

Translatable strings must use a **specific markup** to allow the l10n engine to get the translations.

In a plush template, you can use the `t` helper:

```html
&lt;%= t("greetings") %&gt;
```

### Translation context
<%= sinceVersion("0.10.2") %>

You can use a context with the `t` helper, to provide variables to the translation string:

```html
&lt;%= t("name-format", {name: "Mark"}) %&gt;
```

### Plural handling
<%= sinceVersion("0.10.2") %>

You can use this helper with a numeric second arg to handle singular/plural cases:

```html
&lt;%= t("messages-notification", notificationsLen) %&gt;
```

Provide a context using a third arg:

```html
&lt;%= t("messages-notification", notificationsLen, ctx) %&gt;
```

<%= note() { %>
The second arg is accessible as "Count" in the translations strings.
<% } %>

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

<%= title("Localized Views") %>
<%= sinceVersion("0.10.2") %>

Sometimes, you have to **translate a whole page**, and marking every part of the page takes a lot of time. On some other cases, you'll want to localize the page in a different way for a specific locale. Localized views is a complementary way to handle your translations.

Localized views are **included in the i18n middleware**, so you don't need to setup anything else to use them.

### Create suffixed versions of the templates

First, create a version for the default locale, without a lang suffix:

**page.html**:
```html
&lt;p&gt;This is my default language page.&lt;/p&gt;
```

Then, create a new suffixed version for each language you want to support:

**page.en-us.html**:
```html
&lt;p&gt;This is my en-US version.&lt;/p&gt;
```

**page.fr-fr.html**:
```html
&lt;p&gt;This is my fr-FR version.&lt;/p&gt;
```

The middleware will detect the user language and choose the right template for you! It also works with guest users, using the `Accept-Language` HTTP header. 