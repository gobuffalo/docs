<%= title("Actions") %>

```text
$ buffalo g action --help

Generates new action(s)

Usage:
  buffalo generate action [name] [actionName...] [flags]

Aliases:
  action, a, actions
```

```text
$ buffalo g a users show index create

--> templates/users/show.html
--> templates/users/index.html
--> templates/users/create.html
--> actions/users.go
--> actions/users_test.go
--> goimports -w .
```
