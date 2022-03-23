package godoc

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/docs/search"
	"github.com/gobuffalo/here"
)

type Doc struct {
	here.Info
	Readme string
	Funcs  []Func
}

func (d Doc) Link() string {
	return "https://godoc.org/" + d.ImportPath
}

type Func struct {
	Package string
	Recv    string
	Name    string
	Doc     string
}

func (f Func) String() string {
	bb := &strings.Builder{}
	if len(f.Recv) > 0 {
		bb.WriteString(fmt.Sprintf("(%s) ", f.Recv))
	}
	bb.WriteString(f.Name)
	return bb.String()
}

func (f Func) Link() string {
	if len(f.Recv) > 0 {
		return fmt.Sprintf("https://godoc.org/%s#%s.%s", f.Package, f.Recv, f.Name)
	}
	return fmt.Sprintf("https://godoc.org/%s#%s", f.Package, f.Name)
}

var cache = &docsMap{}

func Indexer(app *buffalo.App) search.Indexer {
	return func() error {
		return Update(app)
	}
}

func Update(app *buffalo.App) error {
	for _, p := range Pkgs {
		go load(p)
	}
	return nil
}

func Get(name string) (*Doc, error) {
	p, ok := cache.Load(name)
	if ok {
		return p, nil
	}
	return load(name)
}

func load(name string) (*Doc, error) {
	info, err := here.Package(name)
	if err != nil {
		return nil, err
	}

	p := &Doc{
		Info:   info,
		Readme: readme(info.Dir),
	}

	for _, f := range info.GoFiles {
		fset := token.NewFileSet()

		node, err := parser.ParseFile(fset, filepath.Join(info.Dir, f), nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		for _, f := range node.Decls {
			fn, ok := f.(*ast.FuncDecl)
			if !ok {
				continue
			}
			df := Func{
				Name:    fn.Name.Name,
				Package: info.ImportPath,
			}
			if fn.Doc != nil {
				df.Doc = fn.Doc.Text()
			}

			if fn.Recv != nil {
				for _, ld := range fn.Recv.List {
					switch t := ld.Type.(type) {
					case *ast.StarExpr:
						if i, ok := t.X.(*ast.Ident); ok {
							df.Recv = i.Name
						}
					case *ast.Ident:
						df.Recv = t.Name
					default:
						continue
					}
				}
			}
			p.Funcs = append(p.Funcs, df)

			d := search.Document{
				URL:    df.Link(),
				Body:   df.Doc,
				Source: search.S_GODOC,
			}
			if err := search.Index(d); err != nil {
				return nil, err
			}
		}
	}
	d := search.Document{
		URL:    p.Link(),
		Body:   p.Readme,
		Source: search.S_GODOC,
	}
	if err := search.Index(d); err != nil {
		return nil, err
	}
	cache.Store(name, p)
	return p, nil
}

func readme(dir string) string {
	b, err := ioutil.ReadFile(filepath.Join(dir, "README.md"))
	if err != nil {
		return ""
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
