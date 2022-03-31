## Append URL and Query String Parameter Duplicates

When multiple query parameters are present with the same name Buffalo would only return the last one when using [`github.com/gobuffalo/buffalo#Context.Params()`](https://godoc.org/github.com/gobuffalo/buffalo#Context.Params()).

### Old Behavior:

```text
GET /users/001?user_id=002&user_id=003
{
    "user_id": [
        "003"
    ]
}
```

### New Behavior:

```text
GET /users/001?user_id=002&user_id=003
{
    "user_id": [
        "001",
        "002",
        "003",
    ]
}
```

* [https://github.com/gobuffalo/buffalo/pull/1778](https://github.com/gobuffalo/buffalo/pull/1778)
