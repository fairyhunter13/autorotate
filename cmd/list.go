package cmd

import (
	"github.com/fairyhunter13/autorotate/sys"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of connected devices",
	Long:  `List of connected devices`,
	Run: func(cmd *cobra.Command, args []string) {
		sys.LogXDisplays()
		sys.LogXInputs()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
