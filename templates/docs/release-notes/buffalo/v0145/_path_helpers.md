## New Path Helpers

`PathFor` takes an `interface{}`, or a `slice` of them, and tries to convert it to a `/foos/{id}` style URL path.

Rules:

* if `string` it is returned as is
* if [`github.com/gobuffalo/helpers/paths#Pathable`](https://godoc.org/github.com/gobuffalo/helpers/paths#Pathable) the `ToPath` method is returned
* if `slice` or an `array` each element is run through the helper then joined
* if [`github.com/gobuffalo/helpers/paths#Paramable`](https://godoc.org/github.com/gobuffalo/helpers/paths#Paramable) the `ToParam` method is used to fill the `{id}` slot
* if `<T>.Slug` the slug is used to fill the `{id}` slot of the URL
* if `<T>.ID` the ID is used to fill the `{id}` slot of the URL
