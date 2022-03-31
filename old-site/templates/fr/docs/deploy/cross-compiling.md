<% seoDescription("Cross-compiler une application Buffalo") %>
<% seoKeywords(["buffalo", "go", "golang", "cross-compilation", "GOOS", "GOARCH", "linux", "arm", "windows", "mac"]) %>

<%= h1("Cross-compiler une application Buffalo") %>

Vous pouvez cross-compiler une application Buffalo, tout comme une application Go normale. Cela signifie qu'il est possible de développer votre application sur un Mac, puis de la compiler pour un serveur Linux depuis votre Mac.

## GOOS et GOARCH

La *toolchain* Go supporte la cross-compilation nativement. Vous devez simplement fournir les variables d'environnement `GOOS` et `GOARCH`.
* `GOOS` définit le système d'exploitation cible (p.e. linux, windows, etc.)
* `GOARCH` définit l'architecture du processeur cible (p.e. amd64, 386, etc.)

Vous pouvez trouver la liste des cibles supportées ici : https://golang.org/doc/install/source#environment (EN)

## Exemples

### Compiler pour Linux AMD64

```go
$ GOOS=linux GOARCH=amd64 buffalo build
```

### Compiler pour Linux ARM64

```go
$ GOOS=linux GOARCH=arm64 buffalo build
```

### Compiler pour Windows i386

```go
$ GOOS=windows GOARCH=386 buffalo build
```