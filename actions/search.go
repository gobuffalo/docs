package actions

import (
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/packr"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/pkg/errors"
)

const indexName = "gobuffalo.search"

var index bleve.Index

func init() {
	os.RemoveAll(indexName)
	var err error
	index, err = bleve.Open(indexName)
	if err == bleve.ErrorIndexPathDoesNotExist {
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexName, mapping)
		if err != nil {
			log.Fatalf("could not create bleve index: %s\n", err)
		}
	}
}

func indexDocs(app *buffalo.App) {
	hl := r.HTMLLayout
	// set a blank layout until we finish indexing
	r.HTMLLayout = ""
	defer func() { r.HTMLLayout = hl }()

	box := r.TemplatesBox
	err := box.Walk(func(path string, file packr.File) error {
		fi, err := file.FileInfo()
		if err != nil {
			return errors.WithStack(err)
		}
		if fi.IsDir() {
			return nil
		}

		if !strings.HasPrefix(path, "docs") {
			return nil
		}

		n := filepath.Base(path)
		if strings.HasPrefix(n, "_") {
			return nil
		}

		u := "/" + path
		ext := filepath.Ext(u)
		for ext != "" {
			u = strings.TrimSuffix(u, ext)
			ext = filepath.Ext(u)
		}

		req := httptest.NewRequest("GET", u, nil)
		req.Header.Set("X-Forwarded-Proto", "https")
		res := httptest.NewRecorder()

		app.ServeHTTP(res, req)
		if res.Code != 200 {
			fmt.Printf("could not index %s\n", u)
			fmt.Println(res.Body.String())
			return nil
		}

		body := strip.StripTags(res.Body.String())
		for strings.Index(body, "  ") > 0 || strings.Index(body, "\n\n") > 0 {
			r := strings.NewReplacer("  ", " ", "\n", " ", "\t", " ")
			body = r.Replace(body)
		}
		doc := struct {
			URL  string
			Body string
		}{
			URL:  u,
			Body: body,
		}

		err = index.Index(doc.URL, doc)
		if err != nil {
			fmt.Println(err)
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func Search(c buffalo.Context) error {
	if c.Param("query") != "" {
		query := bleve.NewQueryStringQuery(c.Param("query"))
		req := bleve.NewSearchRequest(query)
		req.Size = 100
		req.Highlight = bleve.NewHighlight()
		res, err := index.Search(req)
		if err != nil {
			return errors.WithStack(err)
		}
		c.Set("results", res)
	}

	return c.Render(200, r.HTML("search.html"))
}
