package cmd

import (
	"github.com/fairyhunter13/autorotate/sys"
	"github.com/spf13/cobra"
)

var rightCmd = &cobra.Command{
	Use:   "right",
	Short: "Rotate 90deg right.",
	Long:  `Rotate 90deg right.`,
	Run: func(cmd *cobra.Command, args []string) {
		sys.RotateRight(displayName)
	},
}

func init() {
	rootCmd.AddCommand(rightCmd)
}
