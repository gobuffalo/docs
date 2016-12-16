package actions

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"runtime"

	"github.com/markbates/buffalo/render"
	"github.com/markbates/buffalo/render/resolvers"
	"github.com/markbates/gobuffalo/actions/helpers"
)

var r *render.Engine
var resolver = &resolvers.GoPathResolver{Path: "github.com/markbates/gobuffalo"}

func init() {
	r = render.New(render.Options{
		TemplatesPath:  templatesPath(),
		HTMLLayout:     "application.html",
		CacheTemplates: ENV == "production",
	})
	r.RegisterHelper("panel", helpers.PanelHelper)
}

func assetsPath() http.Dir {
	if ENV == "production" {
		return http.Dir("/app/assets")
	}
	p, _ := resolver.Resolve("assets")
	return http.Dir(p)
}

func templatesPath() string {
	if ENV == "production" {
		return "/app/templates"
	}
	p, err := resolver.Resolve("templates")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("### p -> %+v\n", p)
	return p
}

func fromHere(p string) string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), p)
}
