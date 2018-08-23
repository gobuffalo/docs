package actions

import (
	"fmt"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/mmcdole/gofeed"
)

type Video struct {
	Title       string
	Description string
	Link        string
	Published   time.Time
}

var videoList []Video

const videoFeedURL = "https://vimeo.com/channels/gobuffalo/videos/rss"

func indexVideos(app *buffalo.App) {
	app.Logger.Info("Indexing vimeo")

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(videoFeedURL)
	if err != nil {
		app.Logger.Error(err)
		return
	}

	var vl []Video
	for _, item := range feed.Items {
		d := doc{
			URL:  item.Link,
			Body: fmt.Sprintf("[VIDEO] %s", item.Title),
		}

		vl = append(vl, Video{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			Published:   *item.PublishedParsed,
		})

		err := index.Index(d.URL, d)
		if err != nil {
			app.Logger.Error(err)
		}
	}
	videoList = vl

	app.Logger.Infof("%d vimeo items indexed", len(feed.Items))
}
