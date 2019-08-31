package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jyny/outliner/pkg/util"
)

func init() {
	destroyCmd.Flags().StringP("id", "i", "", "ID of instance (required)")
	destroyCmd.Flags().StringP("provider", "p", "", "Provider of instance (required)")
	destroyCmd.MarkFlagRequired("id")
	viper.BindPFlag("id", destroyCmd.Flags().Lookup("id"))
	rootCmd.AddCommand(destroyCmd)
}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy a server",
	Long:  `destroy a server`,
	Run: func(cmd *cobra.Command, args []string) {
		util.PrintDestroyInstanceStart(viper.GetString("id"))
		outliner.DestroyInstance(viper.GetString("id"))
		util.PrintDestroyInstanceDone(viper.GetString("id"))
	},
}
