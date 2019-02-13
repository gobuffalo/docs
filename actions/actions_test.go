package actions_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/gobuffalo/gobuffalo/actions"
	"github.com/gobuffalo/httptest"
	"github.com/stretchr/testify/require"
)

const docsTemplatesDirectory = "../templates/docs/"

func Test_EnDocsHandler(t *testing.T) {
	r := require.New(t)

	w := httptest.New(actions.App())

	files, err := ioutil.ReadDir(docsTemplatesDirectory)

	r.NoError(err)

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			name := strings.Split(file.Name(), ".")[0]
			res := w.HTML("/en/docs/" + name).Get()
			r.Equal(200, res.Code)
		}
	}
}

func Test_FrDocsHandler(t *testing.T) {
	r := require.New(t)

	w := httptest.New(actions.App())

	files, err := ioutil.ReadDir(docsTemplatesDirectory)

	r.NoError(err)

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".fr.md") {
			name := strings.Split(file.Name(), ".")[0]
			res := w.HTML("/fr/docs/" + name).Get()
			r.Equal(200, res.Code)
		}
	}
}
