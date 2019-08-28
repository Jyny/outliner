package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(inspectCmd)
}

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect server",
	Long:  `inspect server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inspect")
	},
}
