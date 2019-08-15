## Routes nommées personnalisées

La fonction [`buffalo.RouteInfo#Name`](https://godoc.org/github.com/gobuffalo/buffalo#RouteInfo.Name) vous permet de donner un nom fixe et personnalisé à un *helper* de route.

```go
a.GET("/coke", CokeHandler).Name("customPath")
```

Cette route est maintenant appelée `customPath`, et vous pouvez y faire référence sous ce nom dans vos templates.

```html
<a href="\<%= customPath() %>">Coke</a>
```

