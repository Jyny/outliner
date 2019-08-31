package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jyny/outliner/pkg/util"
)

func init() {
	deployCmd.Flags().StringP("ip", "i", "", "IP address of Server (required)")
	deployCmd.MarkFlagRequired("ip")
	viper.BindPFlag("ip", deployCmd.Flags().Lookup("ip"))
	rootCmd.AddCommand(deployCmd)
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy outliner to Server",
	Long:  `deploy outliner to Server`,
	PreRun: func(cmd *cobra.Command, args []string) {
		util.PrintDeployInstanceStart()
	},
	Run: func(cmd *cobra.Command, args []string) {
		util.PrintDeployInstanceWait()
		// Todo
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		util.PrintDeployInstanceDone()
	},
}
