package actions

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	gd "github.com/golang/gddo/doc"
	"github.com/pkg/errors"
)

type githubTrans struct{}

func (githubTrans) RoundTrip(req *http.Request) (*http.Response, error) {
	q := req.URL.Query()
	q.Add("access_token", envy.Get("GITHUB_TOKEN", "x"))
	req.URL.RawQuery = q.Encode()
	return http.DefaultTransport.RoundTrip(req)
}

func indexGodocs(app *buffalo.App) error {
	list := []string{
		"github.com/gobuffalo/buffalo",
		"github.com/gobuffalo/pop",
		"github.com/gobuffalo/fizz",
		"github.com/gobuffalo/tags",
		"github.com/gobuffalo/plush",
		"github.com/gobuffalo/packr",
		"github.com/gobuffalo/genny",
		"github.com/gobuffalo/buffalo-plugins",
		"github.com/gobuffalo/buffalo-pop",
		"github.com/gobuffalo/buffalo-goth",
		"github.com/gobuffalo/buffalo-auth",
		"github.com/gobuffalo/buffalo-heroku",
		"github.com/gobuffalo/envy",
		"github.com/gobuffalo/release",
		"github.com/gobuffalo/x",
		"github.com/gobuffalo/flect",
		"github.com/gobuffalo/suite",
		"github.com/gobuffalo/httptest",
		"github.com/gobuffalo/validate",
		"github.com/markbates/grift",
	}
	for _, l := range list {
		if err := indexGodoc(l, app); err != nil {
			// it's fine if there's an error, just log it and move on
			fmt.Println("### err ->", err)
		}
	}
	return nil
}

func indexGodoc(pkg string, app *buffalo.App) error {
	if strings.Contains(pkg, "vendor") {
		return nil
	}
	ctx, cancel := context.WithTimeout(app.Context, 5*time.Second)
	defer cancel()

	c := &http.Client{
		Transport: githubTrans{},
	}
	dpkg, err := gd.Get(ctx, c, pkg, "")
	if err != nil {
		return errors.WithStack(err)
	}
	for _, sd := range dpkg.Subdirectories {
		sd = pkg + "/" + sd
		go func(sd string, app *buffalo.App) {
			err := indexGodoc(sd, app)
			if err != nil {
				fmt.Println(sd, err)
			}
		}(sd, app)
	}

	ub := func(name string) string {
		return "https://godoc.org/" + dpkg.ImportPath + "#" + name
	}

	for _, t := range dpkg.Funcs {
		d := doc{
			URL:  ub(t.Name),
			Body: t.Doc,
		}
		if err := index.Index(d.URL, d); err != nil {
			return errors.WithStack(err)
		}
	}
	for _, t := range dpkg.Types {
		d := doc{
			URL:  ub(t.Name),
			Body: t.Doc,
		}
		if err := index.Index(d.URL, d); err != nil {
			return errors.WithStack(err)
		}
		for _, m := range t.Methods {
			d := doc{
				URL:  ub(t.Name + "." + m.Name),
				Body: m.Doc,
			}
			if err := index.Index(d.URL, d); err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}
