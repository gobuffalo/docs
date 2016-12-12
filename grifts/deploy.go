package grifts

import (
	"os"
	"os/exec"

	. "github.com/markbates/grift/grift"
)

var _ = Add("deploy", func(c *Context) error {
	err := run(exec.Command("git", "push", "origin"))
	if err != nil {
		return err
	}

	err = run(exec.Command("git", "push", "heroku"))
	if err != nil {
		return err
	}

	return run(exec.Command("heroku", "run", "soda", "migrate"))
})

func run(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
