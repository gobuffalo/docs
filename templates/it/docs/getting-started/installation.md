<% seoDescription("Come installare il framework Buffalo") %>
<% seoKeywords(["buffalo", "go", "golang", "installazione", "framework", "web", "mac", "windows", "linux"]) %>

<%= h1("Installa Buffalo") %>

In questo capitolo imparerai a installare Buffalo, sia dai binari precompilati che dai sorgenti.

Buffalo fornisce **due componenti principali**:
* Il tool `buffalo`, un potente toolbox per aiutarti a sviluppare in modo veloce ed efficiente.
* Il frameworl buffalo, un insieme di componenti per costruire la tua app.

Bullafo è attualmente disponibile per le seguenti piattaforme:
* GNU/Linux
* Mac OSX
* Windows

## Dipendenze

Prima di procedere all'installazione assicurati di disporre di quanto segue:

* [Un ambiente Go funzionante](http://gopherguides.com/before-you-come-to-class)
* [Aver incluso `$GOPATH/bin` nella variabile d'ambiente `$PATH`](https://golang.org/doc/code.html#GOPATH)
* [Go](https://golang.org) versione `<%= goMinVersion %>`

##### Dipendenze del Frontend

Le seguenti dipendenze sono opzionali. Non ne hai bisogno se vuoi sviluppare API o se vuoi sviluppare la tua app alla vecchia maniera.

* [node](https://github.com/nodejs/node) versione `8` o successiva
* Uno tra [yarn](https://yarnpkg.com/en/) o [npm](https://github.com/npm/npm) per la [gestione degli asset](/en/docs/assets) che sia predisposto per [webpack](https://github.com/webpack/webpack)

##### Dipendenze specifiche per il database

Anche qui, se non hai bisogno di un database, non hai bisogno di questi software.

* **SQLite 3**: GCC, o un compilatore C equivalente per [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)

## Installazione da un archivio release - 64 bit

<%= note() { %>
I pacchetti release contengo Buffalo senza il supporto a SQLite.
<% } %>

Dalla `v0.10.3`, vengono forniti gli archivi precompilati per ogni release. Se non hai bisogno dell'ultimissima versione, ti suggeriamo di installare questa versione.

### GNU / Linux

```bash
$ wget https://github.com/gobuffalo/buffalo/releases/download/v<%= version %>/buffalo_<%= version %>_Linux_x86_64.tar.gz
$ tar -xvzf buffalo_<%= version %>_Linux_x86_64.tar.gz
$ sudo mv buffalo /usr/local/bin/buffalo
```

### MacOS

```bash
$ curl -OL https://github.com/gobuffalo/buffalo/releases/download/v<%= version %>/buffalo_<%= version %>_Darwin_x86_64.tar.gz
$ tar -xvzf buffalo_<%= version %>_Darwin_x86_64.tar.gz
$ sudo mv buffalo /usr/local/bin/buffalo
# o se hai configurato la directory ~/bin nella variabile d'ambiente PATH
$ mv buffalo ~/bin/buffalo
```

## Scoop (Windows)
Buffalo può essere installato tramite il gestore di pacchetti [Scoop](http://scoop.sh/):

```powershell
PS C:\> scoop install buffalo
```

## Chocolatey (Windows)
Buffalo può essere installato tramite il gestore di pacchetti [Chocolatey](https://chocolatey.org/packages/buffalo). Su Chocolatey le versioni vengono pubblicate con maggior ritardo e devono passare per una moderazione prima di essere rese disponibili:

```powershell
PS C:\> choco install buffalo
```

## Homebrew (macOS)

Su macOS, puoi installare Buffalo anche con [Homebrew](https://brew.sh/). Dopo aver [installato](https://docs.brew.sh/Installation) Homebrew, potrai facilmente installare Buffalo:

```bash
brew install gobuffalo/tap/buffalo
```

## GoFish (Multi-Piattaforma)

[GoFish](https://gofi.sh/index.html) è un gestore di pacchetti multi-piattaforma, che funziona con Windows, MacOSX e Linux.

Dopo aver [installato](https://gofi.sh/index.html#install) GoFish, potrai installare Buffalo davvero facilmente:

```bash
$ gofish install buffalo
==> Installing buffalo...
🐠  buffalo <%= version %>: installed in 3.223672926s
```

## Installazione personalizzata **con** supporto a SQLite3

**SQLite 3** richiede GCC o un compilatore C equivalente per compilare [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3). **Devi** avere un GCC installato **prima** di installare Buffalo.

Per installare Buffalo, assicurati di aver definito `GOPATH`, quindi:

```bash
$ go get -u -v -tags sqlite github.com/gobuffalo/buffalo/buffalo
```

**Utenti Windows**: Segui la guida d'installazione su [https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3](https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3) per installare un GCC su Windows 10. In alternativa, GCC può essere installato con il gestore di pacchetti [Scoop](http://scoop.sh/):

```powershell
PS C:\> scoop install gcc
```

<%= note() { %>
Queste istruzioni possono essere usate anche per aggiornare Buffalo a una nuova versione.
<% } %>

## Installazione personalizzata **senza** supporto a SQLite3

```bash
$ go get -u -v github.com/gobuffalo/buffalo/buffalo
```

<%= note() { %>
Queste istruzioni possono essere usate anche per aggiornare Buffalo a una nuova versione.
<% } %>

## Verifica la tua installazione

Puoi verificare se hai installato tutto in maniera corretta eseguendo il comando `buffalo` in un terminale/prompt dei comandi:

```bash
$ buffalo
Helps you build your Buffalo applications that much easier!

Usage:
  buffalo [command]

Available Commands:
  build       Builds a Buffalo binary, including bundling of assets (packr & webpack)
  db          A tasty treat for all your database needs
  destroy     Allows to destroy generated code.
  dev         Runs your Buffalo app in 'development' mode
  generate    A collection of generators to make life easier
  help        Help about any command
  info        Prints off diagnostic information useful for debugging.
  new         Creates a new Buffalo application
  setup       Setups a newly created, or recently checked out application.
  task        Runs your grift tasks
  test        Runs the tests for your Buffalo app
  update      will attempt to upgrade a Buffalo application to version v<%= version %>
  version     Print the version number of buffalo

Flags:
  -h, --help   help for buffalo

Use "buffalo [command] --help" for more information about a command.
```

Se hai una risposta simile, il tuo toolbox Buffalo è pronto a lavorare!

## Prossimi Passi

* [Integrazione dei Tool](/it/docs/getting-started/integrations) - Lavora con Buffalo, usando i tool esistenti.
* [Creare un Nuovo Progetto](/it/docs/getting-started/new-project) - Crea il tuo primo progetto Buffalo!
