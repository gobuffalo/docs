<%= title("Listing an Application's Middleware", {name: "listing"}) %>

To get a complete list of the middleware your application is using, broken down by grouping, can be found by running the `buffalo task middleware` command.

```bash
$ buffalo t middleware

-> /
github.com/gobuffalo/buffalo.*App.defaultErrorMiddleware
github.com/gobuffalo/buffalo.*App.PanicHandler
github.com/gobuffalo/buffalo.RequestLoggerFunc
github.com/gobuffalo/buffalo.sessionSaver
github.com/gobuffalo/mw-forcessl.Middleware.func1
github.com/markbates/coke/actions.App.func1
github.com/markbates/coke/actions.trackLastURL
github.com/markbates/coke/actions.TrackingCookie
github.com/markbates/coke/actions.App.func3
github.com/gobuffalo/mw-paramlogger.ParameterLogger
github.com/gobuffalo/mw-csrf.glob..func1
github.com/gobuffalo/buffalo-pop/pop/popmw.Transaction.func2
github.com/markbates/coke/actions.SetCurrentUser
github.com/markbates/coke/actions.SetPageTitle
-> /courses/{course_slug}
github.com/gobuffalo/buffalo.*App.defaultErrorMiddleware
github.com/gobuffalo/buffalo.*App.PanicHandler
github.com/gobuffalo/buffalo.RequestLoggerFunc
github.com/gobuffalo/buffalo.sessionSaver
github.com/gobuffalo/mw-forcessl.Middleware.func1
github.com/markbates/coke/actions.App.func1
github.com/markbates/coke/actions.trackLastURL
github.com/markbates/coke/actions.TrackingCookie
github.com/markbates/coke/actions.App.func3
github.com/gobuffalo/mw-paramlogger.ParameterLogger
github.com/gobuffalo/mw-csrf.glob..func1
github.com/gobuffalo/buffalo-pop/pop/popmw.Transaction.func2
github.com/markbates/coke/actions.SetCurrentUser
github.com/markbates/coke/actions.SetPageTitle
github.com/markbates/coke/actions.FindCourse
-> /admin
github.com/gobuffalo/buffalo.*App.defaultErrorMiddleware
github.com/gobuffalo/buffalo.*App.PanicHandler
github.com/gobuffalo/buffalo.RequestLoggerFunc
github.com/gobuffalo/buffalo.sessionSaver
github.com/gobuffalo/mw-forcessl.Middleware.func1
github.com/markbates/coke/actions.App.func1
github.com/markbates/coke/actions.trackLastURL
github.com/markbates/coke/actions.TrackingCookie
github.com/markbates/coke/actions.App.func3
github.com/gobuffalo/mw-paramlogger.ParameterLogger
github.com/gobuffalo/mw-csrf.glob..func1
github.com/gobuffalo/buffalo-pop/pop/popmw.Transaction.func2
github.com/markbates/coke/actions.SetCurrentUser
github.com/markbates/coke/actions.SetPageTitle
github.com/markbates/coke/actions.Authorize
github.com/markbates/coke/actions.AuthorizeAdmin
```
