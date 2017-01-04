package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
)

func Docs(c buffalo.Context) error {
	err := c.Render(200, r.HTML(fmt.Sprintf("docs/%s.md", c.Param("name"))))
	if err != nil {
		return c.Error(404, err)
	}
	return nil
}
