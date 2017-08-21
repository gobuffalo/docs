package helpers

import (
	"html/template"

	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/tags"
	"github.com/markbates/going/randx"
	"github.com/pkg/errors"
	"github.com/shurcooL/github_flavored_markdown"
)

func CodeTabs(help plush.HelperContext) (template.HTML, error) {
	if !help.HasBlock() {
		return "", nil
	}
	s, err := help.Block()
	if err != nil {
		return "", errors.WithStack(err)
	}

	b := github_flavored_markdown.Markdown([]byte(s))

	t := tags.New("div", tags.Options{
		"class": "codetabs",
		"id":    randx.String(10),
	})
	t.Append(tags.New("ul", tags.Options{
		"class": "nav nav-tabs",
		"role":  "tablist",
	}))
	t.Append(tags.New("div", tags.Options{
		"class": "tab-content",
		"body":  string(b),
	}))
	return t.HTML(), nil
}

// <!-- Nav tabs -->
// <ul class="nav nav-tabs" role="tablist">
//   <li role="presentation" class="active"><a href="#home" aria-controls="home" role="tab" data-toggle="tab">Home</a></li>
//   <li role="presentation"><a href="#profile" aria-controls="profile" role="tab" data-toggle="tab">Profile</a></li>
//   <li role="presentation"><a href="#messages" aria-controls="messages" role="tab" data-toggle="tab">Messages</a></li>
//   <li role="presentation"><a href="#settings" aria-controls="settings" role="tab" data-toggle="tab">Settings</a></li>
// </ul>
//
// <!-- Tab panes -->
// <div class="tab-content">
//   <div role="tabpanel" class="tab-pane active" id="home">...</div>
//   <div role="tabpanel" class="tab-pane" id="profile">...</div>
//   <div role="tabpanel" class="tab-pane" id="messages">...</div>
//   <div role="tabpanel" class="tab-pane" id="settings">...</div>
// </div>
