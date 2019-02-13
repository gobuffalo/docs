package actions_test

import (
	"testing"

	"github.com/gobuffalo/gobuffalo/actions"
	"github.com/gobuffalo/httptest"
	"github.com/stretchr/testify/require"
)

func Test_RawHomeHandler(t *testing.T) {
	r := require.New(t)

	w := httptest.New(actions.App())
	res := w.HTML("/").Get()

	r.Equal(302, res.Code)
	r.Equal("/en", res.Location())
}

func Test_HomeHandler(t *testing.T) {
	r := require.New(t)

	w := httptest.New(actions.App())
	res := w.HTML("/en").Get()

	r.Equal(200, res.Code)
}

func Test_HomeHandlerFr(t *testing.T) {
	r := require.New(t)

	w := httptest.New(actions.App())
	res := w.HTML("/fr").Get()

	r.Equal(200, res.Code)
}
