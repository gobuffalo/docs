package helpers

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/plush"
	"github.com/grokify/html-strip-tags-go"
	"github.com/stvp/slug"
)

func init() {
	slug.Replacement = '-'
}

func H1(title string, help plush.HelperContext) template.HTML {
	help.Context.Set("pageTitle", title)
	return template.HTML(fmt.Sprintf("<h1>%s</h1>", title))
}

func SectionTitle(title string, opts map[string]interface{}, help plush.HelperContext) (template.HTML, error) {
	var err error
	if sectionTitleTemplate == nil {
		sectionTitleTemplate, err = plush.Parse(sectionTitleTemplateHTML)
		if err != nil {
			return "", err
		}
	}

	opts["name"] = slug.Clean(strip.StripTags(title))

	if _, ok := opts["title"]; !ok {
		opts["title"] = title
	}

	opts["text"] = title

	s, err := sectionTitleTemplate.Exec(plush.NewContextWith(opts))
	if err != nil {
		return "", err
	}
	return template.HTML(s), err
}

var sectionTitleTemplate *plush.Template

const sectionTitleTemplateHTML = `
<h2>
<a name="<%= name %>" title="<%= htmlEscape(title) %>" href="#<%= name %>"><%= raw(text) %></a>
</h2>
`
