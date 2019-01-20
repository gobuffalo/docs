# Modules Go

**NOTE** : Avant toute chose, veuillez lire [https://github.com/golang/go/wiki/Modules](https://github.com/golang/go/wiki/Modules) pour mieux comprendre comment fonctionnent les modules Go.

<%= title("Activer le support pour les modules Go", {name: "enabling"}) %>

<%= sinceVersion("v0.13.0") %>

Le support pour les modules Go est fourni à titre expérimental dans les paquets fournis par Buffalo, puisque cette fonctionnalité de Go est elle-même encore expérimentale (pour Go `v1.11.x`). Pour activer ce support, vous devez utiliser la variable d'environnement `GO111MODULE` et lui donner la valeur `on`.

C'est un **PRÉ-REQUIS** pour utiliser les modules Go avec Buffalo : la valeur `auto` de la variable `GO111MODULE` n'est **PAS** supportée.

```bash
$ export GO111MODULE=on
```

<%= title("Travailler en dehors du `GOPATH`") %>

En plus de vous permettre de construire vos projets Go à l'identique (avec des versions clairement définies des dépendances), les modules Go offrent la possibilité de travailler facilement en dehors du `GOPATH`.

Avec la configuration `GO111MODULE=on`, la commande `buffalo` devrait fonctionner tout comme elle le faisait avant _dans_ le `GOPATH`.

```bash
$ export GO111MODULE=on
$ buffalo new -h
```

<%= title("Travailler dans le `GOPATH`") %>

Puisque les modules Go sont encore expérimentaux, et incomplets, il est recommandé de continuer à travailler **DANS** le `GOPATH`. Cela vous permettra de tester les modules, tout en pouvant toujours travailler sans.

Lorsque vous travaillez dans le `GOPATH`, vous devriez continuer à utiliser les noms de modules dans le style du `GOPATH`.

#### Recommandé

Ce style de nom de module fonctionne à la fois à l'intérieur et à l'extérieur du `GOPATH`. Il vous permet également de faire fonctionner vos projets avec la commande `go get`.

```go
module github.com/markbates/coke
```

#### Non-recommandé

Ce style de nom de module peut fonctionner dans le `GOPATH`, mais il est moins flexible que le format précédent.

```go
module coke
```

Peu importe le style de nom de module choisi, vous **DEVEZ** respecter le même style dans toute votre application.

Par exemple, si le nom de votre module est `coke`, le nom du paquet contenant les actions est `coke/actions`. Si le nom de votre module est `github.com/markbates/coke`, le nom du paquet contenant les actions est `github.com/markbates/coke/actions`.

<%= title("FAQ") %>

### J'obtiens une erreur `invalid import`

Lorsque le lance `buffalo build`, j'obtiens des erreurs étranges comme celle-ci, **en dehors** de mon `GOPATH`:

```text
invalid import path: "D:/projects/testBuffalo/src/my-project/actions"
```

Assurez-vous que vous avez activé l'option `GO111MODULE=on`. Si ce n'est pas le cas, Buffalo va essayer d'utiliser votre `GOPATH` pour déterminer le chemin de vos paquets. Activez le support pour les modules Go, et réessayez.

### Comment migrer depuis Dep ?

L'outil `go mod init` peut lire votre fichier `Gopkg.toml` et créer un nouveau fichier `go.mod` pour vous. [https://github.com/golang/go/wiki/Modules](https://github.com/golang/go/wiki/Modules)

### Comment utiliser la branche `development` ?

Si vous souhaitez prendre des risques et tester la toute dernière version de Buffalo, vous pouvez demander aux modules Go de récupérer cette version :

```bash
$ go get -u github.com/gobuffalo/buffalo@development
$ go mod tidy
```
