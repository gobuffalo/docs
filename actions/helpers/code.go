package helpers

import (
	"fmt"
	"html"
	"html/template"
	"strings"

	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/tags"
)

func CodeHelper(lang string, opts tags.Options, help plush.HelperContext) (template.HTML, error) {
	classes := []string{fmt.Sprintf("language-%s", lang)}
	if opts["data-line"] != nil {
		classes = append(classes, "line-numbers")
	}
	if opts["class"] != nil {
		classes = append(classes, opts["class"].(string))
	}
	opts["class"] = strings.Join(classes, " ")
	pre := tags.New("pre", opts)
	code := tags.New("code", opts)
	s, err := help.Block()
	if err != nil {
		return "", err
	}
	if lang == "html" {
		s = html.EscapeString(s)
	}
	code.Append(strings.TrimSpace(s))
	pre.Append(code)
	return pre.HTML(), nil
}
