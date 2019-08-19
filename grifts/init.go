package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/docs/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
