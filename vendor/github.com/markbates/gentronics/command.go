package gentronics

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	Should  ShouldFunc
	Command *exec.Cmd
}

func (c *Command) Run(rootPath string, data Data) error {
	if !c.Should(data) {
		return nil
	}
	cmd := c.Command
	fmt.Printf("--> %s\n", strings.Join(cmd.Args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func NewCommand(cmd *exec.Cmd) *Command {
	return &Command{
		Command: cmd,
		Should: func(data Data) bool {
			return true
		},
	}
}
