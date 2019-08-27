package cmd

import (
	"os"
	"fmt"

	_ "github.com/spf13/viper"
	"github.com/spf13/cobra"

	ol "github.com/jyny/outliner/pkg/outliner"
)

func init() {
	cobra.OnInitialize(initConfig)
}
  
func initConfig() {
}

var outliner = ol.New()

var rootCmd = &cobra.Command{
	Use:   "outliner",
	Short: "outliner short",
	Long:  `outliner long`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}