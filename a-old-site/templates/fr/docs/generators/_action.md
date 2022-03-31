## Générer des actions

Puisque créer le templates des actions est plutôt redondant, Buffalo fournit un générateur pour vous aider.

```bash
$ buffalo g action --help

Generates new action(s)

Usage:
  buffalo generate action [name] [actionName...] [flags]

Aliases:
  action, a, actions
```

```bash
$ buffalo g a users show index create

--> templates/users/show.html
--> templates/users/index.html
--> templates/users/create.html
--> actions/users.go
--> actions/users_test.go
--> goimports -w .
```

Il est très probable que vous devrez générer une action pour une méthode HTTP différente de `GET`. Dans ce cas, vous pouvez utiliser le flag `--method`, comme dans l'exemple ci-dessous :

```bash
$ buffalo g actions users message --method POST
```

Dans d'autres scénarios, vous devrez générer une action sans un template HTML associé (pour une API, par exemple). Pour éviter la génération du template HTML lors de la création de l'action, vous pouvez passer le flag `--skip-template` au générateur :

```bash
$ buffalo g actions users update --skip-template
```

<%= note() { %>
C'est le comportement par défaut des applications générées avec le flag `--api`. Consultez [APIs](/fr/docs/apis/) pour plus d'informations.
<% } %>

## Détruire des actions

Vous pouvez supprimer les fichiers générés par ce générateur en exécutant la commande suivante :

```bash
$ buffalo destroy action users
```

Ou dans sa forme courte :

```bash
$ buffalo d a users
```
