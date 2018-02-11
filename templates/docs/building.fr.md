<%= h1("Building Your App") %>

Maintenant, votre projet est prêt à être déployé. Dans cette section vous allez apprendre comment "packager" une version de votre application et la déployée sur un serveur.

<%= title("La commande build") %>

Buffalo contient une command `build`, qui va construire un **éxecutable complet** de votre application, comprenant mais pas seulement; fichiers statiques(i.e ressources graphiques), migrations, templates, etc. Si vous adhérez à la "Manière Buffalo", les choses fonctionnent simplement. C'est une expérience incroyable. :)

```bash
$ buffalo build
```

```bash
Buffalo version <%= version %>

--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o bin/coke -ldflags -X main.version=b5dffda -X main.buildTime="2017-03-20T11:05:23-04:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```

Quand le build est fini, vous avez un binaire tout frais dans le dossier `bin`. Il contiendra aussi **l'heure de compilation** and le **git commit SHA** ancrée, celà permet d'avoir des binaires "versionnés".

<%= title("Personaliser le binaire") %>

Pour avoir la liste des options disponibles, utilsez la commande help:

```bash
$ buffalo help build
```

```bash
Buffalo version <%= version %>

Builds a Buffalo binary, including bundling of assets (packr & webpack)

Usage:
  buffalo build [flags]

Aliases:
  build, b, bill

Flags:
  -c, --compress         compress static files in the binary (default true)
  -e, --extract-assets   extract the assets and put them in a distinct archive
  -h, --help             help for build
      --ldflags string   set any ldflags to be passed to the go build
  -o, --output string    set the name of the binary (default "bin/coke")
  -s, --static           build a static binary using  --ldflags '-linkmode external -extldflags "-static"' (USE FOR CGO)
  -t, --tags string      compile with specific build tags
```

### Nom du binaire / location

Par défaut, votre application aura le binaire dans le dossier `bin` de votre projet, et le nom de l'éxecutable serra le nom que vous avez utilisez pour créez le projet avec la commande `new`.

Vous pouvez changez ce nom par défaut en utilisant le flag `-o` ou `-output`:

```bash
$ buffalo build -o bin/cookies
```

```bash
--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o bin/cookies -ldflags -X main.version="2017-04-02T08:32:28+02:00" -X main.buildTime="2017-04-02T08:32:28+02:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```

En fait, vous pouvez aussi changer le dossier de destination:

```bash
$ # Put the app in my home directory, as "coke"
$ buffalo build -o ~/coke
```

```bash
--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o ~/coke -ldflags -X main.version="2017-04-02T08:32:28+02:00" -X main.buildTime="2017-04-02T08:32:28+02:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```

### Extraire les fichiers statiques dans un fichier Zip.

Par défault, toute l'application est contenue dans un seul exécutable, les fichiers statiques inclus. Dans un système en production, vous voulez peut-être servir ces fichiers statiques avec un serveur proxy (comme Apache ou NGINX), pour réduire la charge de votre application. Vous voudriez peut-être même udiliser un *CDN* pour gérer vos fichiers statiques.

Buffalo fournit un moyen d'extraction des fichiers statiques compilés dans une unique archive, en utilisant le flag `-e` or `-extract-assets`.

```bash
$ buffalo build -e
```

```bash
--> cleaning up target dir
--> running node_modules/.bin/webpack
--> build assets archive
--> disable self assets handling
--> running go build -v -o bin/coke -ldflags -X main.version="2017-04-02T08:45:58+02:00" -X main.buildTime="2017-04-02T08:45:58+02:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```

Par défault, l'archive des fichiers statiques est place dans le répertoire *bin*, mais si vous changez le dossier de destination de l'exécutable avec le flag `-o`, les fichiers statiques seront livrés dans le même répertoire.

```bash
$ ls -la bin
```

```bash
total 36280
drwxr-xr--@  4 markbates  staff   136B Apr  3 10:10 ./
drwxr-xr-x@ 20 markbates  staff   680B Apr  3 10:10 ../
-rwxr-xr-x@  1 markbates  staff    17M Apr  3 10:10 coke*
-rw-r--r--@  1 markbates  staff   691K Apr  3 10:10 coke-assets.zip
```

<%= title("Options avancées") %>

### Construire des binaires "statiques" avec CGO

La création de binaire lié statisquement qui contienent CGO, pensez à SQLite3, peut être délicat. En utilisant le flag `--static` avec `buffalo build`, les flags `--ldflags '-linkmode external -extldflags "-static"'` seront ajoutés à la comande `go build`.

### Build Tags

Quand vous construisez un binnaire Buffalo en utilisant la commande `buffalo build`, vous pouvez passé les flags `--tags` and `--ldflags` pour coustruire le binaire; comme vous le feriez normallement quand vous utilisez l'outils `go build`.

```bash
$ buffalo build --tags="mytag" --ldflags="-X foo.Bar=baz"
```

<%= title("Commandes binaires") %>

### Modes
Le binaire fonctionne par défaut en mode `développement`, ce qui signifie que toutes les sous-commandes seront executées aussi dans ce mode. Pour changer le mode utilisez la variable d'environment `GO_ENV`

```bash
$ GO_ENV=production ./coke
```

### Commandes disponibles

Une fois que le binaire a été construit, il y a plusieurs sous-commandes que vous pouvez éxecuter sur ce binaire:

#### Défault

La commande par défaut, si vous éxecutez simplement le binaire, va démarrer l'application.

#### migrate

La sous-commande `migrate` va éxecuter les commandes de migrations pour l'application.

#### version

La sous-commande `version` va afficher la version du binaire, en y incluant le nom, le SHA du commit git utilisé pour construire le binaire et le moment ou le binaire a été construit.

```bash
$ ./coke version
coke version 69b6a8b ("2017-04-03T10:19:46-04:00")
```

#### task

La sous-commande `task` éxecute les tâches.

```bash
$ ./coke task greet

Hello World!
```
