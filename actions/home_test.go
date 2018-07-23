package actions_test

import (
	"testing"

	"github.com/gobuffalo/gobuffalo/actions"
	"github.com/markbates/willie"
	"github.com/stretchr/testify/require"
)

func Test_RawHomeHandler(t *testing.T) {
	r := require.New(t)

	w := willie.New(actions.App())
	res := w.Request("/").Get()

	r.Equal(302, res.Code)
	r.Equal("/en", res.Location())
}

func Test_HomeHandler(t *testing.T) {
	r := require.New(t)

	w := willie.New(actions.App())
	res := w.Request("/en").Get()

	r.Equal(200, res.Code)
}

func Test_HomeHandlerFr(t *testing.T) {
	r := require.New(t)

	w := willie.New(actions.App())
	res := w.Request("/fr").Get()

	r.Equal(200, res.Code)
}
