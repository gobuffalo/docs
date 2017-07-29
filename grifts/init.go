package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/gobuffalo/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
