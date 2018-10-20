package godoc

import "sort"

func init() {
	sort.Strings(Pkgs)
}

var Pkgs = []string{
	"github.com/gobuffalo/buffalo",
	"github.com/gobuffalo/pop",
	"github.com/gobuffalo/fizz",
	"github.com/gobuffalo/tags",
	"github.com/gobuffalo/plush",
	"github.com/gobuffalo/packr",
	"github.com/gobuffalo/genny",
	"github.com/gobuffalo/buffalo-plugins",
	"github.com/gobuffalo/buffalo-pop",
	"github.com/gobuffalo/buffalo-goth",
	"github.com/gobuffalo/buffalo-auth",
	"github.com/gobuffalo/buffalo-heroku",
	"github.com/gobuffalo/envy",
	"github.com/gobuffalo/events",
	"github.com/gobuffalo/x",
	"github.com/gobuffalo/flect",
	"github.com/gobuffalo/suite",
	"github.com/gobuffalo/httptest",
	"github.com/gobuffalo/validate",
	"github.com/markbates/grift",
}
