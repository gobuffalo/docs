package search

import (
	"github.com/gobuffalo/events"
	"github.com/markbates/oncer"
)

const (
	UNKNOWN int = iota
	S_SITE
	S_BLOG
	S_VIDEO
	S_GODOC

	L_EN
	L_FR
	L_ES

	E_INDEX string = "gobuffalo:search:index"
)

type Indexer func() error

var indexers = []Indexer{}

func AddIndex(i Indexer) {
	indexers = append(indexers, i)
}

func init() {
	oncer.Do("github.com/gobuffalo/search.init", func() {
		events.Listen(func(e events.Event) {
			if e.Kind != E_INDEX {
				return
			}
			for _, i := range indexers {
				go i()
			}
		})
	})
}
