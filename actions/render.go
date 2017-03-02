package actions

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/buffalo/render/resolvers"
	"github.com/gobuffalo/gobuffalo/actions/helpers"
	"github.com/gobuffalo/plush"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		HTMLLayout:     "application.html",
		TemplateEngine: plush.BuffaloRenderer,
		Helpers: map[string]interface{}{
			"panel": helpers.PanelHelper,
		},
		FileResolverFunc: func() resolvers.FileResolver {
			return &resolvers.RiceBox{Box: rice.MustFindBox("../templates")}
		},
	})
}

func assetsPath() http.FileSystem {
	box := rice.MustFindBox("../public/assets")
	return box.HTTPBox()
}
