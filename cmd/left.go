package cmd

import (
	"github.com/fairyhunter13/autorotate/sys"
	"github.com/spf13/cobra"
)

var leftCmd = &cobra.Command{
	Use:   "left",
	Short: "Rotate 90deg left.",
	Long:  `Rotate 90deg left.`,
	Run: func(cmd *cobra.Command, args []string) {
		sys.RotateLeft(displayName)
	},
}

func init() {
	rootCmd.AddCommand(leftCmd)
}
