package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	ol "github.com/jyny/outliner/pkg/outliner"
)

func init() {
	createCmd.Flags().StringP("spec", "s", "", "Spec of Server (required)")
	createCmd.Flags().StringP("region", "r", "", "region of Server (required)")
	createCmd.Flags().StringP("provider", "p", "", "Provider of Server (required)")
	createCmd.MarkFlagRequired("spec")
	createCmd.MarkFlagRequired("region")
	createCmd.MarkFlagRequired("provider")
	viper.BindPFlag("spec", createCmd.Flags().Lookup("spec"))
	viper.BindPFlag("region", createCmd.Flags().Lookup("region"))
	viper.BindPFlag("provider", createCmd.Flags().Lookup("provider"))
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a Server",
	Long:  `create a Server`,
	PreRun: func(cmd *cobra.Command, args []string) {
		printCreateInstanceStart()
	},
	Run: func(cmd *cobra.Command, args []string) {
		inst, err := cloud.CreateInstance(ol.Instance{
			Provider: viper.GetString("provider"),
			SSHKey:   deployer.GetCredentialPub(),
			Region: ol.Region{
				ID: viper.GetString("region"),
			},
			Spec: ol.Spec{
				ID: viper.GetString("spec"),
			},
		})
		if err != nil {
			panic(err)
		}
		printCreateInstanceWait()
		err = cloud.WaitInstance(inst)
		if err != nil {
			panic(err)
		}
		printCreateInstanceDone()
		printInstancesTable([]ol.Instance{inst})
		viper.Set("deploy_ip", inst.IPv4)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		if continueInteractive() {
			deployCmd.PreRun(deployCmd, []string{})
			waitforawhile()
			deployCmd.Run(deployCmd, []string{})
			deployCmd.PostRun(deployCmd, []string{})
		}
	},
}
