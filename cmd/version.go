package cmd

import (
	"github.com/fairyhunter13/autorotate/sys"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Print version`,
	Run: func(cmd *cobra.Command, args []string) {
		isShort, _ := cmd.Flags().GetBool("short")
		sys.LogVersion(isShort)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolP("short", "s", false, "Short version")
}
