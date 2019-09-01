package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	ol "github.com/jyny/outliner/pkg/outliner"
	"github.com/jyny/outliner/pkg/util"
)

func init() {
	inspectCmd.Flags().StringP("id", "i", "", "ID of Server (required)")
	inspectCmd.MarkFlagRequired("id")
	viper.BindPFlag("inspect_id", inspectCmd.Flags().Lookup("id"))
	rootCmd.AddCommand(inspectCmd)
}

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect Server",
	Long:  `inspect Server`,
	Run: func(cmd *cobra.Command, args []string) {
		// id := viper.GetString("inspect_id")
		// util.PrintInstancesTable()
		// util.PrintAPICertTable()
		util.PrintAPICertJSON(ol.APICert{})
	},
}
