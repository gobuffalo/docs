package godoc

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/tags"
	"github.com/pkg/errors"
)

func Helper(help plush.HelperContext) (template.HTML, error) {
	if !help.HasBlock() {
		return "", errors.New("a block is required")
	}

	bb := &bytes.Buffer{}

	for _, pkg := range Pkgs {
		p, err := Get(pkg)
		if err != nil {
			return "", errors.WithStack(err)
		}
		ctx := help.Context.New()
		ctx.Set("pkg", p)
		s, err := help.BlockWith(ctx)
		if err != nil {
			return "", errors.WithStack(err)
		}
		bb.WriteString(s)
	}
	return template.HTML(bb.String()), nil
}

const gdURL = "https://godoc.org/"

func DocLinkHelper(pkg string) template.HTML {
	pkg = strings.TrimPrefix(pkg, gdURL)
	u := gdURL + strings.TrimPrefix(pkg, "*")
	t := tags.New("a", tags.Options{
		"target": "_blank",
		"href":   u,
	})
	c := tags.New("code", tags.Options{})
	c.Append(strings.Replace(pkg, "github.com/gobuffalo/", "", -1))
	t.Append(c)

	return t.HTML()
}
