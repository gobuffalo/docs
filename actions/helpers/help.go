package helpers

import (
	"html/template"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/plush/v4"
	"github.com/pkg/errors"
)

func Faq(title string, opts render.Data, help plush.HelperContext) (template.HTML, error) {
	s, err := help.Block()
	if err != nil {
		return "", errors.WithStack(err)
	}

	name, ok := opts["name"].(string)
	if !ok {
		return "", errors.New("you must supply a name argument for `Faq`")
	}

	ctx := help.Context.New()
	ctx.Set("title", title)
	ctx.Set("body", s)
	ctx.Set("name", name)
	s, err = plush.Render(faqTmpl, ctx)
	return template.HTML(s), err
}

const faqTmpl = `
<li>
	<h6>
		<%= markdown(title) %>
		<a name="<%= name %>"></a>
	</h6>
	<div>
		<%= markdown(body) %>
	</div>
</li>`
