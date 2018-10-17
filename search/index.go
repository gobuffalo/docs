package search

import (
	"os"

	"github.com/blevesearch/bleve"
	"github.com/markbates/oncer"
	"github.com/pkg/errors"
)

const indexName = "gobuffalo.search"

var _index bleve.Index

func index() (bleve.Index, error) {
	var err error
	oncer.Do(indexName, func() {
		os.RemoveAll(indexName)
		var err error
		_index, err = bleve.Open(indexName)
		if err == bleve.ErrorIndexPathDoesNotExist {
			mapping := bleve.NewIndexMapping()
			_index, err = bleve.New(indexName, mapping)
			if err != nil {
				err = errors.WithStack(err)
				return
			}
		}
	})
	return _index, err
}

func Reset() error {
	oncer.Reset(indexName)
	return nil
}

func Index(d Document) error {
	if d.Source == UNKNOWN {
		d.Source = S_SITE
	}
	if d.Language == UNKNOWN {
		d.Language = L_EN
	}
	i, err := index()
	if err != nil {
		return errors.WithStack(err)
	}
	return i.Index(d.URL, d)
}
