package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	ol "github.com/jyny/outliner/pkg/outliner"

	"github.com/jyny/outliner/pkg/cloud/digitalocean"
	"github.com/jyny/outliner/pkg/cloud/linode"
	"github.com/jyny/outliner/pkg/cloud/vultr"
)

// Persistent Flags
var apikeycfg string
var id_rsa string
var id_rsa_pub string

// Persi	stent outliner for other commends
var outliner = ol.New()

var rootCmd = &cobra.Command{
	Use:   "outliner",
	Short: "outliner short",
	Long:  `outliner long`,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&apikeycfg, "apikey", "k", "", "export api key from file")
}

func initConfig() {
	if apikeycfg != "" {
		viper.SetConfigFile(apikeycfg)
	}

	// add cloud providers
	outliner.ActivateProvider(
		digitalocean.Activator{},
		linode.Activator{},
		vultr.Activator{},
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
