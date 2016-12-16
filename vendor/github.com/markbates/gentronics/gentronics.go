package gentronics

import (
	"os"

	"github.com/pkg/errors"
)

type Data map[string]interface{}

type RunFn func(string, Data) error

type Runnable interface {
	Run(string, Data) error
}

type Generator struct {
	Runners []Runnable
}

func New() *Generator {
	return &Generator{
		Runners: []Runnable{},
	}
}

func (g *Generator) Add(r Runnable) {
	g.Runners = append(g.Runners, r)
}

func (g *Generator) Run(rootPath string, data Data) error {
	err := os.MkdirAll(rootPath, 0755)
	if err != nil {
		return errors.WithStack(err)
	}
	err = os.Chdir(rootPath)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, r := range g.Runners {
		err := r.Run(rootPath, data)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
