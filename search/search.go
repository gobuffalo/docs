package search

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
	go i()
	indexers = append(indexers, i)
}
