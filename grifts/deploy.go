package grifts

import (
	"os"
	"os/exec"

	. "github.com/markbates/grift/grift"
)

var _ = Add("deploy", func(c *Context) error {
	cmd := exec.Command("git", "push", "heroku")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("heroku", "run", "soda", "migrate")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
})
