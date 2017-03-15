package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/buffalo/render/resolvers"
	"github.com/gobuffalo/gobuffalo/actions/helpers"
	"github.com/gobuffalo/packr"
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
			return &resolvers.PackrBox{
				Box: packr.NewBox("../templates"),
			}
		},
	})
}

func assetsPath() http.FileSystem {
	box := packr.NewBox("../public/assets")
	return box.HTTPBox()
}
