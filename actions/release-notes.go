package actions

import (
	"html/template"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packd"
	"github.com/google/go-cmp/cmp"
)

type releaseNotes struct {
	box packd.Box
}

<<<<<<< HEAD
func (releaseNotes) Path(key string, v string, ps ...string) string {
=======
func (releaseNotes) Path(key, v string, ps ...string) string {
>>>>>>> master
	args := []string{
		"docs",
		"release-notes",
		key,
		strings.ReplaceAll(v, ".", ""),
	}
	args = append(args, ps...)
	return filepath.Join(args...)
}

func (rn releaseNotes) Diff(key, c1, c2 string, ps ...string) (template.HTML, error) {
	p1 := rn.Path(key, c1, ps...)
	p2 := rn.Path(key, c2, ps...)

	b1, err := rn.box.Find(p1)
	if err != nil {
		return "", err
	}

	b2, err := rn.box.Find(p2)
	if err != nil {
		return "", err
	}

	d := cmp.Diff(string(b1), string(b2))

	return template.HTML(d), nil
}
