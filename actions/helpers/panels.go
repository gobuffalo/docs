package helpers

import (
	"html/template"

	"github.com/gobuffalo/plush"
	"github.com/markbates/inflect"
	"github.com/shurcooL/github_flavored_markdown"
)

func PanelHelper(title string, opts map[string]interface{}, help plush.HelperContext) (template.HTML, error) {
	var err error
	if panelTemplate == nil {
		panelTemplate, err = plush.Parse(panelTemplateHtml)
		if err != nil {
			return "", err
		}
	}

	body, err := help.Block()
	if err != nil {
		return "", err
	}

	opts["body"] = template.HTML(github_flavored_markdown.Markdown([]byte(body)))
	if _, ok := opts["style"]; !ok {
		opts["style"] = "primary"
	}
	if _, ok := opts["name"]; !ok {
		opts["name"] = inflect.Dasherize(title)
	}

	opts["title"] = title
	s, err := panelTemplate.Exec(plush.NewContextWith(opts))
	if err != nil {
		return "", err
	}
	return template.HTML(s), err
}

var panelTemplate *plush.Template

const panelTemplateHtml = `
<a name="<%= name %>" title="<%= htmlEscape(title) %>"></a>
<div class="panel panel-<%= style %>">
  <div class="panel-heading">
    <h3><%= title %></h3>
  </div>
  <div class="panel-body">
	<%= body %>
  </div>
</div>`
