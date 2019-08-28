package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all servers",
	Long:  `list all servers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%+v", outliner.ListInstance())
	},
}
