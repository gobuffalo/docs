package helpers

import (
	"html/template"

	"github.com/gobuffalo/plush"
	"github.com/markbates/inflect"
)

func SectionTitle(title string, opts map[string]interface{}, help plush.HelperContext) (template.HTML, error) {
	var err error
	if sectionTitleTemplate == nil {
		sectionTitleTemplate, err = plush.Parse(sectionTitleTemplateHtml)
		if err != nil {
			return "", err
		}
	}

	if _, ok := opts["name"]; !ok {
		opts["name"] = inflect.Dasherize(title)
	}

	opts["title"] = title
	s, err := sectionTitleTemplate.Exec(plush.NewContextWith(opts))
	if err != nil {
		return "", err
	}
	return template.HTML(s), err
}

var sectionTitleTemplate *plush.Template

const sectionTitleTemplateHtml = `
<h2>
<a name="<%= name %>" title="<%= htmlEscape(title) %>" href="#<%= name %>"><%= title %></a>
</h2>
`
