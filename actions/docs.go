package actions

import (
	"fmt"

	"github.com/markbates/buffalo"
)

func Docs(c buffalo.Context) error {
	return c.Render(200, r.HTML(fmt.Sprintf("docs/%s.md", c.Param("name"))))
}
