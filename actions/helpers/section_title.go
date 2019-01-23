package helpers

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/plush"
	"github.com/stvp/slug"
)

func init() {
	slug.Replacement = '-'
}

func H1(title string, help plush.HelperContext) template.HTML {
	help.Context.Set("pageTitle", title)
	return template.HTML(fmt.Sprintf("<h1>%s</h1>", title))
}

var sectionTitleTemplate *plush.Template

const sectionTitleTemplateHTML = `
<h2>
<a name="<%= name %>" title="<%= htmlEscape(title) %>" href="#<%= name %>"><%= raw(text) %></a>
</h2>
`
