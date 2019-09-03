package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jyny/outliner/pkg/util"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show outliner version",
	Long:  `show outliner version`,
	Run: func(cmd *cobra.Command, args []string) {
		util.PrintVersion(version)
	},
}
