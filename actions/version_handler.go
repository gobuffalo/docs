package actions

import (
	"errors"
	"os"

	"github.com/markbates/buffalo"
	"github.com/markbates/gobuffalo/models"
	"github.com/markbates/pop"
)

func VersionList(c buffalo.Context) error {
	v := []models.BuffaloVersion{}
	tx := c.Get("tx").(*pop.Connection)
	err := tx.All(&v)
	if err != nil {
		return c.Error(500, err)
	}
	return c.Render(200, r.JSON(v))
}

func VersionCurrent(c buffalo.Context) error {
	v := models.CurrentBuffaloVersion(c.Get("tx").(*pop.Connection))
	return c.Render(200, r.JSON(v))
}

func VersionUpdate(c buffalo.Context) error {
	v := &models.BuffaloVersion{}
	err := c.Bind(v)
	if err != nil {
		return err
	}
	vt := os.Getenv("VERSION_TOKEN")
	if v.Token != vt || vt == "" {
		return c.Error(422, errors.New("bad token!"))
	}

	tx := c.Get("tx").(*pop.Connection)
	err = tx.Create(v)
	if err != nil {
		return err
	}
	return c.Render(201, r.JSON(v))
}

func SetVersion(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		v := models.CurrentBuffaloVersion(c.Get("tx").(*pop.Connection))
		c.Set("version", v.Version)
		return next(c)
	}
}
