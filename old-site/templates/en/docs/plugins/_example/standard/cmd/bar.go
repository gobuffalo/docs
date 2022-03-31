package cmd

import (
	"github.com/spf13/cobra"
)

// barCmd represents the bar command
var barCmd = &cobra.Command{
	Use:   "bar",
	Short: "description about this plugin",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(barCmd)
}
