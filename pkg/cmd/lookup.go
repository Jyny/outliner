package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	lookupCmd.AddCommand(specCmd)
	lookupCmd.AddCommand(regineCmd)
	rootCmd.AddCommand(lookupCmd)
}

var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "lookup available regiens",
	Long:  `lookup available regiens`,
}

var regineCmd = &cobra.Command{
	Use:   "regien",
	Short: "lookup available regiens",
	Long:  `lookup available regiens`,
	Run: func(cmd *cobra.Command, args []string) {
		printRegions(outliner.LookupRegion())
	},
}

var specCmd = &cobra.Command{
	Use:   "spec",
	Short: "lookup available regiens",
	Long:  `lookup available regiens`,
	Run: func(cmd *cobra.Command, args []string) {
		printSpecs(outliner.LookupSpec())
	},
}
