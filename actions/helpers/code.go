package helpers

import (
	"bytes"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/github_flavored_markdown"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/tags"
	"github.com/markbates/going/randx"
	"github.com/pkg/errors"
)

func CodeTabs(help plush.HelperContext) (template.HTML, error) {
	if !help.HasBlock() {
		return "", nil
	}
	s, err := help.Block()
	if err != nil {
		return "", errors.WithStack(err)
	}
	return codeTab(s)
}

func codeTab(s string) (template.HTML, error) {
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

func ExampleDir(r *render.Engine) func(dir string, help plush.HelperContext) (template.HTML, error) {
	return func(dir string, help plush.HelperContext) (template.HTML, error) {
		files := map[string][]genny.File{}

		box := r.TemplatesBox
		err := box.WalkPrefix(dir, func(path string, f packr.File) error {
			if filepath.Base(path) == ".DS_Store" {
				return nil
			}
			name := strings.TrimPrefix(path, dir+string(filepath.Separator))

			key := "/"
			dirs := strings.Split(name, string(filepath.Separator))
			if len(dirs) > 1 {
				key = dirs[0]
			}

			gfs, ok := files[key]
			if !ok {
				gfs = []genny.File{}
			}
			gfs = append(gfs, genny.NewFile(name, f))
			files[key] = gfs

			return nil
		})
		if err != nil {
			return "", errors.WithStack(err)
		}

		div := tags.New("div", tags.Options{})
		for k, gfs := range files {
			dt := tags.New("div", tags.Options{})
			title := strings.TrimSpace(k)
			if !strings.HasPrefix(title, "/") {
				title = "/" + title
			}
			title = "." + title
			dt.Append(tags.New("h3", tags.Options{"body": title}))

			bb := &bytes.Buffer{}
			for _, f := range gfs {
				ext := strings.TrimPrefix(filepath.Ext(f.Name()), ".")
				if len(ext) == 0 {
					ext = "plain"
				}
				bb.WriteString("\n```")
				bb.WriteString(ext + "\n")
				bb.WriteString("// " + f.Name() + "\n")
				bb.WriteString(f.String())
				bb.WriteString("\n```\n")
			}

			ct, err := codeTab(bb.String())
			if err != nil {
				return "", errors.WithStack(err)
			}
			dt.Append(ct)
			div.Append(dt)
		}

		return div.HTML(), nil
	}
}
