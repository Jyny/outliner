package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path"

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
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	// define `.env`
	viper.SetConfigType("env")
	viper.SetConfigName("")

	// search from possible paths
	viper.AddConfigPath(path.Join(u.HomeDir, "/.outliner/"))
	viper.AddConfigPath(u.HomeDir)
	viper.AddConfigPath(".")

	// top precedence order is explicit config file path
	viper.SetConfigFile(apikeycfg)

	err = viper.ReadInConfig()
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
