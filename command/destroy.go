package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"

	"github.com/jyny/outliner/pkg/util"
)

func init() {
	destroyCmd.Flags().StringP("id", "i", "", "ID of Server (required)")
	destroyCmd.MarkFlagRequired("id")
	viper.BindPFlag("destroy_id", destroyCmd.Flags().Lookup("id"))
	RootCmd.AddCommand(destroyCmd)
}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy a Server",
	Long:  `destroy a Server`,
	PreRun: func(cmd *cobra.Command, args []string) {
		util.PrintDestroyInstanceStart()
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := viper.GetString("destroy_id")
		log.Println(id)
		err := cloud.DestroyInstance(id)
		if err != nil {
			panic(err)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		util.PrintDestroyInstanceDone()
	},
}
