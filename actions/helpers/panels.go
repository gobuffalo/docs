package helpers

import (
	"html/template"

	"github.com/gobuffalo/velvet"
	"github.com/markbates/inflect"
)

func PanelHelper(help velvet.HelperContext) (template.HTML, error) {
	var err error
	if panelTemplate == nil {
		panelTemplate, err = velvet.Parse(panelTemplateHtml)
		if err != nil {
			return "", err
		}
	}

	body, err := help.Block()
	if err != nil {
		return "", err
	}

	data := help.Context.New()
	data.Set("body", template.HTML(body))

	if !data.Has("style") {
		data.Set("style", "primary")
	}
	if !data.Has("name") {
		data.Set("name", inflect.Dasherize(data.Get("title").(string)))
	}
	s, err := panelTemplate.Exec(data)
	if err != nil {
		return "", err
	}
	return template.HTML(s), err
}

var panelTemplate *velvet.Template

const panelTemplateHtml = `
<a name="{{name}}" data-title="{{html_escape title}}"></a>
<div class="panel panel-{{style}}">
  <div class="panel-heading">
    <h3 class="panel-title">{{title}}</h3>
  </div>
  <div class="panel-body">
	{{body}}
  </div>
</div>`
