package actions_test

import (
	"os"
	"testing"

	"github.com/gobuffalo/gobuffalo/actions"
	"github.com/gobuffalo/gobuffalo/models"
	"github.com/markbates/pop"
	"github.com/markbates/willie"
	"github.com/stretchr/testify/require"
)

func Test_VersionHandler(t *testing.T) {
	w := willie.New(actions.App())
	os.Setenv("VERSION_TOKEN", "12345")

	t.Run("valid token", func(st *testing.T) {
		tx(func(tx *pop.Connection) {
			r := require.New(st)
			v := &models.BuffaloVersion{
				Version: "1.0.0",
				Token:   "12345",
			}
			res := w.JSON("/version").Post(v)
			r.Equal(201, res.Code)

			err := tx.First(v)
			r.NoError(err)
			r.NotEmpty(v.ID)
		})
	})
	t.Run("invalid token", func(st *testing.T) {
		tx(func(tx *pop.Connection) {
			r := require.New(st)
			v := &models.BuffaloVersion{
				Version: "2.0.0",
				Token:   "bad",
			}
			res := w.JSON("/version").Post(v)
			r.Equal(422, res.Code)

			err := tx.Where("version = ?", "2.0.0").First(v)
			r.Error(err)
		})
	})
}
