package helpers

import (
	"fmt"

	"github.com/aymerick/raymond"
)

func PanelHelper(options *raymond.Options) raymond.SafeString {
	var err error
	if panelTemplate == nil {
		panelTemplate, err = raymond.Parse(panelTemplateHtml)
		if err != nil {
			return raymond.SafeString(fmt.Sprintf("<pre><code>%s</pre></code>", err.Error()))
		}
	}
	data := options.Hash()
	data["body"] = raymond.SafeString(options.Fn())
	if _, ok := data["style"]; !ok {
		data["style"] = "primary"
	}
	s, err := panelTemplate.Exec(data)
	if err != nil {
		return raymond.SafeString(fmt.Sprintf("<pre><code>%s</pre></code>", err.Error()))
	}
	return raymond.SafeString(s)
}

var panelTemplate *raymond.Template

const panelTemplateHtml = `<div class="panel panel-{{style}}">
  <div class="panel-heading">
    <h3 class="panel-title">{{title}}</h3>
  </div>
  <div class="panel-body">
	{{body}}
  </div>
</div>`
