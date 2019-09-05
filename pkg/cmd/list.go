package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jyny/outliner/pkg/util"
)

func init() {
	listCmd.AddCommand(infoCmd)
	listCmd.AddCommand(instCmd)
	listCmd.AddCommand(specCmd)
	listCmd.AddCommand(regineCmd)
	listCmd.AddCommand(providerCmd)
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list following [command]",
	Long:  `list following things`,
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "list information needed by create",
	Long:  `list information needed by create`,
	Run: func(cmd *cobra.Command, args []string) {
		specCmd.Run(specCmd, []string{})
		regineCmd.Run(regineCmd, []string{})
	},
}

var instCmd = &cobra.Command{
	Use:   "server",
	Short: "list active Servers",
	Long:  `list active Servers`,
	Run: func(cmd *cobra.Command, args []string) {
		insts, err := cloud.ListInstance()
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
		regs, err := cloud.ListRegion()
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
		specs, err := cloud.ListSpec()
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
		pvdrs, err := cloud.ListProvider()
		if err != nil {
			panic(err)
		}
		util.PrintProvidersTable(pvdrs)
	},
}
