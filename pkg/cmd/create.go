package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	ol "github.com/jyny/outliner/pkg/outliner"
	"github.com/jyny/outliner/pkg/util"
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
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a Server",
	Long:  `create a Server`,
	PreRun: func(cmd *cobra.Command, args []string) {
		util.PrintCreateInstanceStart()
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
		util.PrintCreateInstanceWait()
		err = cloud.WaitInstance(inst)
		if err != nil {
			panic(err)
		}
		util.PrintCreateInstanceDone()
		util.PrintInstancesTable([]ol.Instance{inst})
		viper.Set("ip", inst.IPv4)
		viper.Set("id", inst.ID)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		if util.ContinueInteractive() {
			deployCmd.PreRun(deployCmd, []string{})
			util.Waitforawhile()
			deployCmd.Run(deployCmd, []string{})
			deployCmd.PostRun(deployCmd, []string{})
		}
	},
}
