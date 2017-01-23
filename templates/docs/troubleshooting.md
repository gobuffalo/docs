# Trouble Shooting

{{#panel title="buffalo new <project> fails looking for golang.org/x/tools/go/gcimporter"}}

This is caused by an outdated copy of the `github.com/motemen/gore` package. To fix simply update `gore`:

```text
$ go get -u github.com/motemen/gore
```

For information see [https://github.com/gobuffalo/buffalo/issues/108](https://github.com/gobuffalo/buffalo/issues/108) and [https://github.com/motemen/gore/issues/63](https://github.com/motemen/gore/issues/63).
{{/panel}}

{{#panel title="buffalo dev fails to start"}}

When starting `$ buffalo dev`, and you encounter this error:

`There was a problem starting the dev server: Unknown, Please review the troubleshooting docs`

This may be due to your system missing NodeJS/NPM, Ensure that Node/NPM is installed and is in your `$PATH`. If  Node/NPM are indeed in your `$PATH`, try renaming webpack.config.js. 

If you are still having issues after attempting the steps above, please reach out to the community in the #buffalo channel on Gophers Slack.