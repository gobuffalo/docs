package actions

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/gobuffalo/search/blog"
	"github.com/pkg/errors"
)

// HomeHandler serves the home page.
func HomeHandler(c buffalo.Context) error {
	c.Set("blogPosts", blog.LastPosts)
	return c.Render(200, r.HTML("overview.html", "home.html"))
}

// Sponsors serves the sponsors page.
func Sponsors(c buffalo.Context) error {
	return c.Render(200, r.HTML("sponsors.html"))
}

// ChangeLanguage handles the selector to change the current locale.
func ChangeLanguage(c buffalo.Context) error {
	f := struct {
		OldLanguage string `form:"oldLanguage"`
		Language    string `form:"language"`
		URL         string `form:"url"`
	}{}
	if err := c.Bind(&f); err != nil {
		return errors.WithStack(err)
	}

	// Set new language prefix
	if f.URL == fmt.Sprintf("/%s", f.OldLanguage) {
		f.URL = ""
	} else {
		p := fmt.Sprintf("/%s/", f.OldLanguage)
		f.URL = strings.TrimPrefix(f.URL, p)
	}

	return c.Redirect(302, fmt.Sprintf("/%s/%s", f.Language, f.URL))
}
