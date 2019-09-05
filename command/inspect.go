package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	ol "github.com/jyny/outliner/pkg/outliner"
	"github.com/jyny/outliner/pkg/util"
)

func init() {
	inspectCmd.Flags().StringP("id", "i", "", "ID of Server (required)")
	inspectCmd.MarkFlagRequired("id")
	viper.BindPFlag("id", inspectCmd.Flags().Lookup("id"))
	RootCmd.AddCommand(inspectCmd)
}

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect Server",
	Long:  `inspect Server`,
	Run: func(cmd *cobra.Command, args []string) {
		id := viper.GetString("id")
		inst, err := cloud.InspectInstance(id)
		if err != nil {
			panic(err)
		}
		apicert, err := deployer.GetServiceCert(inst.IPv4)
		if err != nil {
			panic(err)
		}
		util.PrintInstancesTable([]ol.Instance{inst})
		util.PrintAPICertTable(apicert)
		util.PrintAPICertJSON(apicert)
	},
}
