# Console/REPL

**_Currently only available in `tip` or versions greater than `0.4.6`._**

Buffalo wraps the [https://github.com/motemen/gore](https://github.com/motemen/gore) REPL and pre-loads it with the `actions` and `models` packages from your application, should they exist.

{{#panel title="Installation"}}

Currently it is not possible to set the `gore` package as a dependency in Buffalo, because of that you must first install `gore` before you can use the console.

```text
$ go get -u github.com/motemen/gore
```
{{/panel}}

{{#panel title="Using the Console" name="using"}}

```text
$ buffalo console
# or
$ buffalo c
```
<img src="/assets/images/repl.gif" width="100%">
{{/panel}}
