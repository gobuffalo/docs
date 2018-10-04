package actions

import (
	"context"
	"fmt"
	"net/http"
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
	}
	for _, l := range list {
		if err := indexGodoc(l, app); err != nil {
			fmt.Println("### err ->", err)
		}
	}
	return nil
}

func indexGodoc(pkg string, app *buffalo.App) error {
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
