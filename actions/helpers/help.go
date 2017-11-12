package helpers

import (
	"html/template"

	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

func Faq(title string, help plush.HelperContext) (template.HTML, error) {
	s, err := help.Block()
	if err != nil {
		return "", errors.WithStack(err)
	}
	ctx := help.Context.New()
	ctx.Set("title", title)
	ctx.Set("body", s)
	s, err = plush.Render(faqTmpl, ctx)
	return template.HTML(s), err
}

const faqTmpl = `
<li>
	<h6><%= markdown(title) %></h6>
	<div>
		<%= markdown(body) %>
	</div>
</li>`
