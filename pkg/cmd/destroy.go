package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jyny/outliner/pkg/util"
)

func init() {
	destroyCmd.Flags().StringP("id", "i", "", "ID of Server (required)")
	destroyCmd.MarkFlagRequired("id")
	viper.BindPFlag("id", destroyCmd.Flags().Lookup("id"))
	rootCmd.AddCommand(destroyCmd)
}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy a Server",
	Long:  `destroy a Server`,
	PreRun: func(cmd *cobra.Command, args []string) {
		util.PrintDestroyInstanceStart()
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := viper.GetString("id")
		err := cloud.DestroyInstance(id)
		if err != nil {
			panic(err)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		util.PrintDestroyInstanceDone()
	},
}
