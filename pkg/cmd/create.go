package cmd

import (
	"github.com/spf13/cobra"

	ol "github.com/jyny/outliner/pkg/outliner"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a server",
	Long:  `create a server`,
	Run: func(cmd *cobra.Command, args []string) {
		outliner.CreateInstance(ol.Instance{
			Region: ol.Region{
				ID: "",
			},
			Spec: ol.Spec{
				ID: "",
			},
		})
	},
}
