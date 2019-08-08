package blog

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/gobuffalo/search"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
)

func Indexer(app *buffalo.App) search.Indexer {
	return func() error {
		return index(app)
	}
}

func index(app *buffalo.App) error {
	res, err := http.Get(FeedURL)
	if err != nil {
		app.Logger.Error(err)
		return errors.WithStack(err)
	}
	if res.StatusCode != 200 {
		err := errors.Errorf("failed to index blog %d", res.StatusCode)
		app.Logger.Error(err)
		return errors.WithStack(err)
	}

	blog := &Feed{}
	err = json.NewDecoder(res.Body).Decode(blog)
	if err != nil {
		app.Logger.Error(err)
		return errors.WithStack(err)
	}

	for _, b := range blog.Items {
		body := strip.StripTags(b.Content)
		for strings.Index(body, "  ") > 0 || strings.Index(body, "\n\n") > 0 {
			r := strings.NewReplacer("  ", " ", "\n", " ", "\t", " ")
			body = r.Replace(body)
		}
		d := search.Document{
			URL:    b.Link,
			Body:   body,
			Source: search.S_BLOG,
		}
		err = search.Index(d)
		if err != nil {
			app.Logger.Error(err)
			return errors.WithStack(err)
		}
	}

	truncateString := func(str string, num int) string {
		bnoden := []rune(str)
		if len(bnoden) > num {
			if num > 3 {
				num -= 3
			}
			bnoden = append(bnoden[0:num], '.', '.', '.')
		}
		return string(bnoden)
	}

	p := bluemonday.StrictPolicy()
	for i, bp := range blog.Items[:3] {
		ip := strings.Index(bp.Description, "</p>")
		if ip != -1 {
			bp.Description = bp.Description[:ip]
		}
		bp.Description = truncateString(p.Sanitize(bp.Description), 143)
		mu.Lock()
		lastPosts[i] = bp
		mu.Unlock()
	}

	return nil
}
