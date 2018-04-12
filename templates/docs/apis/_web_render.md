```go
func init() {
	r = render.New(render.Options{
		HTMLLayout:   "application.html",
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,
		Helpers:      render.Helpers{},
	})
}
```
