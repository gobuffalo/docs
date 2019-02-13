package site

import (
	"fmt"
	"net/http/httptest"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/events"
	"github.com/gobuffalo/gobuffalo/search"
	"github.com/gobuffalo/packr/v2"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/pkg/errors"
)

type indexer struct {
	App    *buffalo.App
	Engine *render.Engine
}

func Indexer(app *buffalo.App, r *render.Engine) search.Indexer {
	return indexer{
		App:    app,
		Engine: r,
	}.Index

}

func (ind indexer) Index() error {
	r := ind.Engine
	defer events.EmitPayload("gobuffalo:site:search:finished", events.Payload{})

	box := r.TemplatesBox
	err := box.Walk(func(path string, file packr.File) error {
		fi, err := file.FileInfo()
		if err != nil {
			return errors.WithStack(err)
		}
		if fi.IsDir() {
			return nil
		}

		if !strings.HasPrefix(path, "docs/") {
			return nil
		}

		for _, n := range strings.Split(path, string(filepath.Separator)) {
			if strings.HasPrefix(n, "_") {
				return nil
			}
		}

		u := "/en/" + path
		ext := filepath.Ext(u)
		for ext != "" {
			u = strings.TrimSuffix(u, ext)
			ext = filepath.Ext(u)
		}

		req := httptest.NewRequest("GET", u, nil)
		req.Header.Set("X-Forwarded-Proto", "https")
		q := req.URL.Query()
		q.Add("_indexing", "true")
		req.URL.RawQuery = q.Encode()
		res := httptest.NewRecorder()

		ind.App.ServeHTTP(res, req)
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
		d := search.Document{
			URL:      u,
			Body:     body,
			Source:   search.S_SITE,
			Language: search.L_EN,
		}

		return search.Index(d)
	})
	if err != nil {
		ind.App.Logger.Error(err)
		return err
	}
	return nil
}
