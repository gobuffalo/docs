---
name: Configuration
seoDescription: "Handle configuration with Buffalo"
seoKeywords:
  - "buffalo"
  - "go"
  - "golang"
  - "configuration"
  - "env"
  - "framework"
  - "web"
weight: 4
aliases:
  - /pt/documentation/getting_started/config-vars
  - /pt/docs/getting-started/config-vars
  - /pt/docs/getting-started/config-vars
---

# Configuration

In this chapter, you'll learn how to manage configuration with Buffalo.

Environment variables are a good way to separate environment specific values, or secrets, from your application code base ([as described in The Twelve Factor app](https://12factor.net/config)). It can help define behavior that is based on the context of the app (as requiring SSL on production) and isolate secrets from your code base. This way, all developers don't have to know the productions keys to sensitive services, such as a bank API, and they can use sandbox API keys.

## Available Environment Variables

The following variables are used by Buffalo:

| Variable                 | Default                  | Usage                                                      |
| ---                      | ---                      | ---                                                        |
| `GO_ENV`                 | `development`            | The environment (dev, qa, production etc.) that Buffalo is run in.                   |
| `GO_BIN`                 | `go`                     | The Go compiler to use.                                     |
| `BUFFALO_PLUGIN_PATH`    | `$PATH`                  | Where Buffalo looks for plugins.                            |
| `BUFFALO_PLUGIN_TIMEOUT` | `1s`                     | How long Buffalo waits for a plugin to respond.             |
| `ADDR`                   | `127.0.0.1` or `0.0.0.0` | Which address to bind the server to.                        |
| `PORT`                   | `3000`                   | Which port to bind the server to.                           |
| `HOST`                   | `http://127.0.0.1:$PORT` | The "URL" of the application (i.e. what end users type in). |
| `SESSION_SECRET`         | `""`                     | A salt used for securing sessions.                          |

## Custom Configuration

You still can provide your own variables, and retrieve them from within your application. The [envy](https://github.com/gobuffalo/envy) package makes it easy!

```go
import "github.com/gobuffalo/envy"

// [...]

// Get MYSECRET env variable, default to empty string if it's not set
var MYSECRET = envy.Get("MYSECRET", "")

// Get REQUIREDSECRET env variable, return an error if it's not set
REQUIREDSECRET, err := envy.MustGet("REQUIREDSECRET")
```

## Support for .env Files

<!--%= sinceVersion("0.10.3") %-->

Buffalo ships with `.env` support (**since buffalo >= 0.10.3**), meaning buffalo will load `.env` files into environment variables once the application starts. To do it, Buffalo uses [`envy.Load`](https://github.com/gobuffalo/envy/blob/e613c80275b86293880eddeb27417c9a7c670ff3/envy.go#L53) which will look for `.env` file at the root of your app.

If you're not familiar with how a `.env` file looks, here is an example:

```text
SENDGRID_API_KEY=ABCCOQ7GFRVCW0ODHPFQ3FTP5SLL1Q
SENDGRID_EMAIL=email@myapp.com

APP_DEBUG=true
APP_LOG_LEVEL=debug
APP_URL=https://myapp.com
```

Generated apps (**with buffalo >= 0.10.3**) will also create a default `.env` file in your application root. This file will be watched by Buffalo for changes, but will be ignored by git (added in the `.gitignore`). This is a good way to prevent developers to push credentials by mistake.

## Next Steps

* [Tooling Integration](/documentation/getting_started/integrations) - Work with Buffalo, using existing tools.
