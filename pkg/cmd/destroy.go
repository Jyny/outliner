package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(destroyCmd)
}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy a server",
	Long:  `destroy a server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("destroy")
	},
}
