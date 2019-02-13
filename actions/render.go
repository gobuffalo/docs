package actions

import (
	"fmt"
	"html/template"
	"path"
	"strings"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/gobuffalo/actions/helpers"
	"github.com/gobuffalo/gobuffalo/search/godoc"
	"github.com/gobuffalo/gobuffalo/search/vimeo"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/markbates/inflect"
)

var r *render.Engine
var assetBox = packr.New("app:assets", "../public")

func Renderer() *render.Engine {
	return r
}

func init() {
	r = render.New(render.Options{
		HTMLLayout: "application.html",
		Helpers: render.Helpers{
			"doclink":   godoc.DocLinkHelper,
			"goDocPkgs": godoc.Pkgs,
			"godoc":     godoc.Helper,
			"h1":        helpers.H1,
			"note":      helpers.Note,
			"warning":   helpers.Warning,
			"sinceVersion": func(version string, opts render.Data, help plush.HelperContext) (template.HTML, error) {
				ctx := help.Context.New()
				if !strings.HasPrefix(version, "v") {
					version = "v" + version
				}
				var name string
				pkg := "github.com/gobuffalo/buffalo"
				if n, ok := opts["pkg"].(string); ok {
					pkg = n
					name = strings.TrimPrefix(n, "github.com/") + " "
				}
				ctx.Set("name", name)
				ctx.Set("version", version)
				ctx.Set("pkg", pkg)
				s, err := plush.Render(sinceVersion, ctx)
				return template.HTML(s), err
			},
			"vimeo": func(code string) template.HTML {
				return template.HTML(fmt.Sprintf(vimeoTmpl, code))
			},
			"vimeoFromVideo": func(v vimeo.Video) template.HTML {
				code := path.Base(v.Link)
				return template.HTML(fmt.Sprintf(vimeoTmpl, code))
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
		TemplatesBox: packr.New("../templates", "../templates"),
		AssetsBox:    assetBox,
	})
	r.Helpers["exampleDir"] = helpers.ExampleDir(r)
}

const vimeoTmpl = `<div class="video">
<iframe src="https://player.vimeo.com/video/%s?portrait=0" width="640" height="360" frameborder="0" webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>
</div>`

const sinceVersion = `<span class="since-version"><a href="https://<%= pkg %>/releases/tag/<%= version %>" target="_blank"><%= raw(t("helpers.since", {"version": version, "name": name})) %></a></span>`

const githubRelease = `<a href="https://github.com/gobuffalo/buffalo/releases/tag/%s" target="_blank" rel="noopener noreferrer">releases/tag/%s</a>`
