package cmd

import (
	"fmt"
	"os"
	"os/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	ol "github.com/jyny/outliner/pkg/outliner"

	"github.com/jyny/outliner/pkg/cloud/digitalocean"
	"github.com/jyny/outliner/pkg/cloud/linode"
	"github.com/jyny/outliner/pkg/cloud/vultr"
)

// Persistent Flags
var apikeycfg string
var sshkey string
var sshkeyPub string

// Persistent outliner for other commends
var outliner = ol.New()

var rootCmd = &cobra.Command{
	Use:   "outliner",
	Short: "outliner short",
	Long:  `outliner long`,
}

// Execute entry of commandline
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&apikeycfg, "env", "E", "", "export api key from file")
}

func initConfig() {
	if apikeycfg != "" {
		viper.SetConfigFile(apikeycfg)
	} else {
		usr, err := user.Current()
		if err != nil {
			panic(err)
		}
		viper.AddConfigPath(usr.HomeDir)
		viper.SetConfigName("outliner.env")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Activate & register cloud providers
	outliner.RegisterProvider(
		validater,
		digitalocean.Activator{},
		linode.Activator{},
		vultr.Activator{},
	)
}
