package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lookupCmd)
}

var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "lookup available regiens",
	Long:  `lookup available regiens`,
	Run: func(cmd *cobra.Command, args []string) {
		printRegions(outliner.LookupRegion())
	},
}
