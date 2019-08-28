package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a server",
	Long:  `create a server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create")
	},
}
