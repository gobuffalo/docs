package actions

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/gobuffalo/actions/helpers"
	"github.com/gobuffalo/packr"
)

var r *render.Engine
var assetBox = packr.NewBox("../public/assets")

func init() {
	r = render.New(render.Options{
		HTMLLayout: "application.html",
		Helpers: render.Helpers{
			"h1":    helpers.H1,
			"title": helpers.SectionTitle,
			"note":  helpers.Note,
			"sinceVersion": func(version string) template.HTML {
				return template.HTML(fmt.Sprintf(sinceVersion, version))
			},
			"vimeo": func(code string) template.HTML {
				return template.HTML(fmt.Sprintf(vimeo, code))
			},
			"codeTabs": helpers.CodeTabs,
			"faq":      helpers.Faq,
			"githubRelease": func(version string) template.HTML {
				return template.HTML(fmt.Sprintf(githubRelease, version, version))
			},
		},
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetBox,
	})
}

const vimeo = `<div class="video">
<iframe src="https://player.vimeo.com/video/%s?portrait=0" width="640" height="360" frameborder="0" webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>
</div>`

const sinceVersion = `<span class="label label-primary">since <strong>v%s</strong></span>`

const githubRelease = `<a href="https://github.com/gobuffalo/buffalo/releases/tag/%s" target="_blank" rel="noopener noreferrer">releases/tag/%s</a>`
