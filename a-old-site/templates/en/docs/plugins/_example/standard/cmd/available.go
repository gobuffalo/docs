package cmd

import (
	"encoding/json"
	"os"

	"github.com/gobuffalo/buffalo-plugins/plugins"
	"github.com/spf13/cobra"
)

// availableCmd represents the available command
var availableCmd = &cobra.Command{
	Use:   "available",
	Short: "a list of available buffalo plugins",
	RunE: func(cmd *cobra.Command, args []string) error {
		plugs := plugins.Commands{
			// {Name: "bar", UseCommand: "generate", BuffaloCommand: "generate", Description: generateCmd.Short, Aliases: generateCmd.Aliases},
		}
		return json.NewEncoder(os.Stdout).Encode(plugs)
	},
}

func init() {
	rootCmd.AddCommand(availableCmd)
}
