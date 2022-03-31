package cmd

import (
	"fmt"

	"github.com/foo/buffalo-bar/bar"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "current version of bar",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("bar", bar.Version)
		return nil
	},
}

func init() {
	barCmd.AddCommand(versionCmd)
}
