# Getting Started

![Buffalo Logo](/assets/images/logo_med.png)

Buffalo is Go web framework. Yeah, I hate the word "framework" too! Buffalo is different though. Buffalo doesn't want to re-invent wheels like routing and templating. Buffalo is glue that wraps all of the best packages available and makes them all play nicely together.

Buffalo is "idiomatic", for whatever that is worth. The purpose of a framework is not to bend you to it's will, but to rather get out of your way and make your job of building your application easy. That is the goal of Buffalo.

If you were to look through the Buffalo code base you'll find little code, just enough to assemble the amazing packages that other's have written into one coherent system.

> I :heart: web dev in go again - Brian Ketelsen

## Installation

```
$ go get -u github.com/markbates/buffalo/buffalo
```

## Generating a new Project

Buffalo aims to make building new web applications in Go as simple as possible, and what could be more simple that a new application generator?

```
$ buffalo new <name>
```

That will generate a whole new Buffalo application that is ready to go. It'll even run `go get` for you to make sure you have all of the necessary dependencies needed to run your application.

To see a list of available flags for the `new` command, just check out it's help.

```
$ buffalo help new
```

## Running your application

Buffalo is Go "standards" compliant, that means you can just build your binary and run it. It's that simple.

### Running your application in Development

One of the downsides to Go development is the lack of code "reloading". This means as you change your code you need to manually stop your application, rebuild it, and then restart it. Buffalo finds this is annoying, and wants to make life better for you.

```
$ buffalo dev
```

The `dev` command will watch your `.go` and `.html` files by default and rebuild, and restart, your binary for you so you don't have to worry about such things. Just run the `dev` command and start coding.
