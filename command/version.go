package command

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show outliner version",
	Long:  `show outliner version`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersion(version)
	},
}
