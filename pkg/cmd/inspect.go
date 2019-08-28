package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	inspectCmd.Flags().StringP("id", "i", "", "ID of instance (required)")
	inspectCmd.Flags().StringP("provider", "p", "", "Provider of instance (required)")
	inspectCmd.MarkFlagRequired("id")
	viper.BindPFlag("id", destroyCmd.Flags().Lookup("id"))
	rootCmd.AddCommand(inspectCmd)
}

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect server",
	Long:  `inspect server`,
	Run: func(cmd *cobra.Command, args []string) {
		id := viper.GetString("id")
		_, err := outliner.InspectInstance(id)
		if err != nil {
			panic(err)
		}
	},
}
