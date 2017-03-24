package actions

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/gobuffalo/actions/helpers"
	"github.com/gobuffalo/packr"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		HTMLLayout: "application.html",
		Helpers: render.Helpers{
			"title": helpers.SectionTitle,
			"code":  helpers.CodeHelper,
		},
		TemplatesBox: packr.NewBox("../templates"),
	})
}
