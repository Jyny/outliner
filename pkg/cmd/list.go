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
		insts, err := outliner.ListInstance()
		if err != nil {
			panic(err)
		}
		util.PrintInstancesTable(insts)
	},
}

var regineCmd = &cobra.Command{
	Use:   "regien",
	Short: "list available regiens",
	Long:  `list available regiens`,
	Run: func(cmd *cobra.Command, args []string) {
		regs, err := outliner.ListRegion()
		if err != nil {
			panic(err)
		}
		util.PrintRegionsTable(regs)
	},
}

var specCmd = &cobra.Command{
	Use:   "spec",
	Short: "list available specs",
	Long:  `list available specs`,
	Run: func(cmd *cobra.Command, args []string) {
		specs, err := outliner.ListSpec()
		if err != nil {
			panic(err)
		}
		util.PrintSpecsTable(specs)
	},
}

var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "list available providers",
	Long:  `list available providers`,
	Run: func(cmd *cobra.Command, args []string) {
		pvdrs, err := outliner.ListProvider()
		if err != nil {
			panic(err)
		}
		util.PrintProvidersTable(pvdrs)
	},
}
