package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jyny/outliner/pkg/util"
)

func init() {
	listCmd.AddCommand(instCmd)
	listCmd.AddCommand(specCmd)
	listCmd.AddCommand(regineCmd)
	listCmd.AddCommand(providerCmd)
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list following [command]",
	Long:  `list following things`,
}

var instCmd = &cobra.Command{
	Use:   "instance",
	Short: "list active instances(servers)",
	Long:  `list active instances(servers)`,
	Run: func(cmd *cobra.Command, args []string) {
		util.PrintListInstances(outliner.ListInstance())
	},
}

var regineCmd = &cobra.Command{
	Use:   "regien",
	Short: "list available regiens",
	Long:  `list available regiens`,
	Run: func(cmd *cobra.Command, args []string) {
		util.PrintRegions(outliner.ListRegion())
	},
}

var specCmd = &cobra.Command{
	Use:   "spec",
	Short: "list available specs",
	Long:  `list available specs`,
	Run: func(cmd *cobra.Command, args []string) {
		util.PrintSpecs(outliner.ListSpec())
	},
}

var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "list available providers",
	Long:  `list available providers`,
	Run: func(cmd *cobra.Command, args []string) {
		util.PrintProvider(outliner.ListProvider())
	},
}
