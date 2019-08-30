package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
	Use:   "server",
	Short: "list active servers(instance)",
	Long:  `list active servers(instance)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%+v", outliner.ListInstance())
	},
}

var regineCmd = &cobra.Command{
	Use:   "regien",
	Short: "list available regiens",
	Long:  `list available regiens`,
	Run: func(cmd *cobra.Command, args []string) {
		printRegions(outliner.ListRegion())
	},
}

var specCmd = &cobra.Command{
	Use:   "spec",
	Short: "list available specs",
	Long:  `list available specs`,
	Run: func(cmd *cobra.Command, args []string) {
		printSpecs(outliner.ListSpec())
	},
}

var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "list available providers",
	Long:  `list available providers`,
	Run: func(cmd *cobra.Command, args []string) {
		printProvider(outliner.ListProvider())
	},
}
