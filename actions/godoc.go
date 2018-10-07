package actions

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/plush"
	gd "github.com/golang/gddo/doc"
	"github.com/pkg/errors"
)

func init() {
	sort.Strings(goDocPkgs)
}

var godoc = &GoDoc{
	moot: &sync.RWMutex{},
	data: map[string]*gd.Package{},
}

var goDocPkgs = []string{
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
	"github.com/gobuffalo/x",
	"github.com/gobuffalo/flect",
	"github.com/gobuffalo/suite",
	"github.com/gobuffalo/httptest",
	"github.com/gobuffalo/validate",
	"github.com/markbates/grift",
}

type githubTrans struct{}

func (githubTrans) RoundTrip(req *http.Request) (*http.Response, error) {
	q := req.URL.Query()
	q.Add("access_token", envy.Get("GITHUB_TOKEN", "x"))
	req.URL.RawQuery = q.Encode()
	return http.DefaultTransport.RoundTrip(req)
}

type GoDoc struct {
	moot  *sync.RWMutex
	data  map[string]*gd.Package
	udata map[string]*gd.Package
}

func (g *GoDoc) List() []*gd.Package {
	g.moot.RLock()
	var list []*gd.Package
	for _, v := range g.data {
		list = append(list, v)
	}
	g.moot.RUnlock()
	sort.Slice(list, func(a, b int) bool {
		return list[a].Name < list[b].Name
	})
	return list
}

func (g *GoDoc) Get(pkg string) (*gd.Package, error) {
	g.moot.RLock()
	if p, ok := g.data[pkg]; ok {
		g.moot.RUnlock()
		return p, nil
	}
	g.moot.RUnlock()

	p := &gd.Package{
		ProjectName: pkg,
		ImportPath:  pkg,
		ProjectURL:  "https://" + pkg,
	}

	return p, nil
}

func (g *GoDoc) Update(ctx context.Context) error {
	g.moot.Lock()
	g.udata = map[string]*gd.Package{}
	g.moot.Unlock()
	for _, l := range goDocPkgs {
		if err := g.indexGodoc(ctx, l); err != nil {
			// it's fine if there's an error, just log it and move on
			fmt.Println("### err ->", err)
		}
	}
	g.moot.Lock()
	if len(g.udata) != 0 {
		g.data, g.udata = g.udata, g.data
	}
	g.moot.Unlock()
	return nil
}

func (g *GoDoc) indexGodoc(ctx context.Context, pkg string) error {
	if strings.Contains(pkg, "vendor") {
		return nil
	}
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)

	c := &http.Client{
		Transport: githubTrans{},
	}
	dpkg, err := gd.Get(ctx, c, pkg, "")
	if err != nil {
		return errors.WithStack(err)
	}
	g.moot.Lock()
	g.udata[pkg] = dpkg
	g.moot.Unlock()
	for _, sd := range dpkg.Subdirectories {
		sd = pkg + "/" + sd
		go func(sd string, ctx context.Context) {
			err := g.indexGodoc(ctx, sd)
			if err != nil {
				fmt.Println(sd, err)
			}
		}(sd, ctx)
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

func indexGodocs(app *buffalo.App) error {
	return godoc.Update(app.Context)
}

func godocHelper(help plush.HelperContext) (template.HTML, error) {
	if !help.HasBlock() {
		return "", errors.New("a block is required")
	}

	bb := &bytes.Buffer{}

	for _, pkg := range goDocPkgs {
		p, err := godoc.Get(pkg)
		if err != nil {
			return "", errors.WithStack(err)
		}
		ctx := help.Context.New()
		ctx.Set("pkg", p)
		s, err := help.BlockWith(ctx)
		if err != nil {
			return "", errors.WithStack(err)
		}
		bb.WriteString(s)
	}
	return template.HTML(bb.String()), nil
}
