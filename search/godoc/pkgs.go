package godoc

import "sort"

func init() {
	sort.Strings(Pkgs)
}

var Pkgs = []string{
	"github.com/gobuffalo/buffalo",
	"github.com/gobuffalo/pop",
	"github.com/gobuffalo/fizz",
	"github.com/gobuffalo/tags/v3",
	"github.com/gobuffalo/plush/v4",
	"github.com/gobuffalo/packr/v2",
	"github.com/gobuffalo/genny",
	"github.com/gobuffalo/buffalo-pop",
	"github.com/gobuffalo/buffalo-goth",
	"github.com/gobuffalo/buffalo-auth",
	"github.com/gobuffalo/buffalo-heroku",
	"github.com/gobuffalo/envy",
	"github.com/gobuffalo/meta",
	"github.com/gobuffalo/logger",
	"github.com/gobuffalo/nulls",
	"github.com/gobuffalo/attrs",
	"github.com/gobuffalo/licenser",
	"github.com/gobuffalo/mapi",
	"github.com/gobuffalo/mw-csrf",
	"github.com/gobuffalo/mw-i18n",
	"github.com/gobuffalo/mw-basicauth",
	"github.com/gobuffalo/mw-contenttype",
	"github.com/gobuffalo/mw-forcessl",
	"github.com/gobuffalo/mw-paramlogger",
	"github.com/gobuffalo/mw-tokenauth",
	"github.com/gobuffalo/mw-poptx",
	"github.com/gobuffalo/events",
	"github.com/gobuffalo/x",
	"github.com/gobuffalo/flect",
	"github.com/gobuffalo/suite",
	"github.com/gobuffalo/syncx",
	"github.com/gobuffalo/shoulders",
	"github.com/gobuffalo/httptest",
	"github.com/gobuffalo/validate",
	"github.com/gobuffalo/clara",
	"github.com/gobuffalo/release",
	"github.com/gobuffalo/lush",
	"github.com/markbates/grift",
	"github.com/markbates/pkger",
}
