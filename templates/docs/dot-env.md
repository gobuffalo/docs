# .env Support

<%= sinceVersion("0.10.3") %>

Environment variables are a good way to separate environment specific values or secrets from your application codebase ([as described in The Twelve Factor app](https://12factor.net/config)), it can help define behavior that is based in the context of the app and isolate secrets from your codebase so p.e all developers doesn't have to know the productions key to bank account API and they can use locally development API keys.

If you're not familiar with how a .env file looks, here is an example:

```
SENDGRID_API_KEY=ABCCOQ7GFRVCW0ODHPFQ3FTP5SLL1Q
SENDGRID_EMAIL=email@myapp.com

APP_DEBUG=true
APP_LOG_LEVEL=debug
APP_URL=https://myapp.com
```

Buffalo ships with `.env` support (**since buffalo >= 0.10.3**), meaning buffalo will load .env files into ENV variables once the application starts, to do it, buffalo uses [`Envy.Load`](https://github.com/gobuffalo/envy/blob/e613c80275b86293880eddeb27417c9a7c670ff3/envy.go#L53) which will look for `.env` file at the root of your app.

Generated apps (**with buffalo >= 0.10.3**) will also create a default `.env` file in your application root, this file will be listened by refresh for changes but will be ignored by git (added in the gitignore).

