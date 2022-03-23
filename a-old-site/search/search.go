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

	E_INDEX string = "docs:search:index"
)

type Indexer func() error
