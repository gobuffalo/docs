package actions

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/gobuffalo/actions/helpers"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/markbates/inflect"
)

var r *render.Engine
var assetBox = packr.NewBox("../public")

func init() {
	r = render.New(render.Options{
		HTMLLayout: "application.html",
		Helpers: render.Helpers{
			"goDocPkgs": goDocPkgs,
			"godoc":     godocHelper,
			"h1":        helpers.H1,
			"title":     helpers.SectionTitle,
			"note":      helpers.Note,
			"warning":   helpers.Warning,
			"sinceVersion": func(version string, help plush.HelperContext) (template.HTML, error) {
				ctx := help.Context.New()
				ctx.Set("version", version)
				s, err := plush.Render(sinceVersion, ctx)
				return template.HTML(s), err
			},
			"vimeo": func(code string) template.HTML {
				return template.HTML(fmt.Sprintf(vimeo, code))
			},
			"codeTabs": helpers.CodeTabs,
			"faq":      helpers.Faq,
			"githubRelease": func(version string) template.HTML {
				return template.HTML(fmt.Sprintf(githubRelease, version, version))
			},
			"seoDescription": func(description string, help plush.HelperContext) {
				help.Context.Set("metaDescription", description)
			},
			"seoKeywords": func(keywords []interface{}, help plush.HelperContext) {
				kl := make([]string, 0)
				for _, k := range keywords {
					kl = append(kl, k.(string))
				}
				help.Context.Set("metaKeywords", strings.Join(kl, ","))
			},
			"pluralize_uint": func(s string, i uint64) string {
				return inflect.PluralizeWithSize(s, int(i))
			},
			"score": func(f float64) string {
				return fmt.Sprintf("%.2f%%", f*100)
			},
		},
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetBox,
	})
	r.Helpers["exampleDir"] = helpers.ExampleDir(r)
}

const vimeo = `<div class="video">
<iframe src="https://player.vimeo.com/video/%s?portrait=0" width="640" height="360" frameborder="0" webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>
</div>`

const sinceVersion = `<span class="since-version"><%= raw(t("helpers.since", {"version": version})) %></span>`

const githubRelease = `<a href="https://github.com/gobuffalo/buffalo/releases/tag/%s" target="_blank" rel="noopener noreferrer">releases/tag/%s</a>`
