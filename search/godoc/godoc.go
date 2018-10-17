package godoc

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/gobuffalo/search"
	gd "github.com/golang/gddo/doc"
	"github.com/pkg/errors"
)

var godoc = &GoDoc{
	moot: &sync.RWMutex{},
	data: map[string]*gd.Package{},
}

const tk = "GITHUB_TOKEN"

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
	g.add(p)

	return p, nil
}

func (g *GoDoc) add(p *gd.Package) {
	g.moot.Lock()
	p.ProjectName = strings.TrimPrefix(p.ImportPath, "github.com/")
	p.Doc = readme(p)
	g.data[p.ImportPath] = p
	g.moot.Unlock()
}

func readme(p *gd.Package) string {
	if len(p.Doc) > 0 {
		return p.Doc
	}
	if !strings.HasPrefix(p.ImportPath, "github.com") {
		return ""
	}
	u := "https://raw.githubusercontent.com/" + p.ProjectName + "/master/README.md"
	res, err := http.Get(u)
	if err != nil {
		return p.Doc
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return p.Doc
	}
	s := string(b)

	bb := &bytes.Buffer{}

	var h1 bool
	for _, line := range strings.Split(s, "\n") {
		if strings.HasPrefix(line, "#") {
			h1 = true
			continue // don't write the h1
		}
		if h1 {
			bb.WriteString(line + "\n")
			continue
		}
	}

	return strings.TrimSpace(bb.String())
}

func (g *GoDoc) Update(ctx context.Context) error {
	g.moot.Lock()
	g.udata = map[string]*gd.Package{}
	g.moot.Unlock()
	for _, l := range Pkgs {
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
	g.add(dpkg)
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
		d := search.Document{
			URL:    ub(t.Name),
			Body:   t.Doc,
			Source: search.S_GODOC,
		}
		if err := search.Index(d); err != nil {
			return errors.WithStack(err)
		}
	}
	for _, t := range dpkg.Types {
		d := search.Document{
			URL:    ub(t.Name),
			Body:   t.Doc,
			Source: search.S_GODOC,
		}
		if err := search.Index(d); err != nil {
			return errors.WithStack(err)
		}
		for _, m := range t.Methods {
			d := search.Document{
				URL:    ub(t.Name + "." + m.Name),
				Body:   m.Doc,
				Source: search.S_GODOC,
			}
			if err := search.Index(d); err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}

func Indexer(app *buffalo.App) search.Indexer {
	return func() error {
		return godoc.Update(app.Context)
	}
}
