package gentronics

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type ShouldFunc func(Data) bool

type File struct {
	Path       string
	Template   string
	Permission os.FileMode
	Should     ShouldFunc
}

func (f *File) Run(rootPath string, data Data) error {
	if !f.Should(data) {
		return nil
	}

	path, err := f.render(f.Path, data)
	if err != nil {
		return err
	}

	body, err := f.render(f.Template, data)
	if err != nil {
		return err
	}

	return f.save(rootPath, path, body)
}

func (f *File) save(rootPath, path, body string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(filepath.Join(rootPath, dir), f.Permission)
	if err != nil {
		return err
	}

	odir := filepath.Join(rootPath, path)
	fmt.Printf("--> %s\n", odir)

	ff, err := os.Create(odir)
	if err != nil {
		return err
	}

	_, err = ff.WriteString(body)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (f *File) render(s string, data Data) (string, error) {
	t, err := template.New("").Parse(s)
	if err != nil {
		return "", err
	}
	bb := bytes.Buffer{}
	err = t.Execute(&bb, data)
	return bb.String(), err
}

func NewFile(path string, template string) *File {
	return &File{
		Path:       path,
		Template:   template,
		Permission: 0755,
		Should: func(data Data) bool {
			return true
		},
	}
}
