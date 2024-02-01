package cmd

import (
	"github.com/fairyhunter13/autorotate/sys"
	"github.com/spf13/cobra"
)

var portraitCmd = &cobra.Command{
	Use:   "portrait",
	Short: "Portrait screen orientation.",
	Long:  `Portrait screen orientation.`,
	Run: func(cmd *cobra.Command, args []string) {
		isInvert, _ := cmd.Flags().GetBool("invert")

		if isInvert {
			sys.RotateRight(displayName)
		} else {
			sys.RotateLeft(displayName)
		}
	},
}

func init() {
	rootCmd.AddCommand(portraitCmd)

	portraitCmd.Flags().BoolP("invert", "i", false, "Invert portrait orientation.")
}
