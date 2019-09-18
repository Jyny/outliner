package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	ol "github.com/jyny/outliner/pkg/outliner"
)

func init() {
	inspectCmd.Flags().StringP("ip", "i", "", "IP address of Server (required)")
	inspectCmd.MarkFlagRequired("ip")
	viper.BindPFlag("inspect_ip", inspectCmd.Flags().Lookup("ip"))
	RootCmd.AddCommand(inspectCmd)
}

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect Server",
	Long:  `inspect Server`,
	PreRun: func(cmd *cobra.Command, args []string) {
		ip := viper.GetString("inspect_ip")
		inst, err := cloud.InspectInstanceByIP(ip)
		if err == nil {
			printInstancesTable([]ol.Instance{inst})
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		ip := viper.GetString("inspect_ip")
		apicert, err := deployer.GetServiceCert(ip)
		if err != nil {
			panic(err)
		}
		printAPICertTable(apicert)
		printAPICertJSON(apicert)
	},
}
