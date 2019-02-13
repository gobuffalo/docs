package vimeo

import (
	"fmt"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/gobuffalo/search"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

type Video struct {
	Title       string
	Description string
	Link        string
	Published   time.Time
}

var Videos []Video

const FeedURL = "https://vimeo.com/channels/gobuffalo/videos/rss"

func Indexer(app *buffalo.App) search.Indexer {
	return func() error {
		return indexVideos(app)
	}
}

func indexVideos(app *buffalo.App) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(FeedURL)
	if err != nil {
		app.Logger.Error(err)
		return errors.WithStack(err)
	}

	var vl []Video
	for _, item := range feed.Items {
		d := search.Document{
			URL:    item.Link,
			Body:   fmt.Sprintf("[VIDEO] %s", item.Title),
			Source: search.S_VIDEO,
		}

		vl = append(vl, Video{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			Published:   *item.PublishedParsed,
		})

		err := search.Index(d)
		if err != nil {
			app.Logger.Error(err)
			return errors.WithStack(err)
		}
	}
	Videos = vl
	return nil
}
