package godoc

import (
	"bytes"
	"html/template"

	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

func Helper(help plush.HelperContext) (template.HTML, error) {
	if !help.HasBlock() {
		return "", errors.New("a block is required")
	}

	bb := &bytes.Buffer{}

	for _, pkg := range Pkgs {
		p, err := godoc.Get(pkg)
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
