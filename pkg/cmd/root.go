package cmd

import (
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	//"github.com/jyny/outliner/pkg/cloud/digitalocean"
	"github.com/jyny/outliner/pkg/cloud/linode"
	//"github.com/jyny/outliner/pkg/cloud/vultr"

	ol "github.com/jyny/outliner/pkg/outliner"
	"github.com/jyny/outliner/pkg/util"
)

// Persistent Flags
var cfgFile string

// Persistent outliner for other commends
var outliner = ol.New()

var rootCmd = &cobra.Command{
	Use:   "outliner",
	Short: "Auto setup & deploy tool for outline VPN server",
	Long:  "Auto setup & deploy tool for outline VPN server",
}

// Execute entry of commandline
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "F", "", "config file (default is $HOME/.outliner/.env)")
}

func initConfig() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	// `.env` as config file name
	viper.SetConfigType("env")
	viper.SetConfigName("")

	// search from possible paths
	viper.AddConfigPath(filepath.Join(u.HomeDir, "/.outliner/"))
	viper.AddConfigPath(u.HomeDir)
	viper.AddConfigPath(".")

	if cfgFile != "" {
		// top precedence order of paths
		viper.SetConfigFile(cfgFile)
	} else {
		// set flag to load config from $ENV
		viper.AutomaticEnv()
	}

	// load config file
	viper.ReadInConfig()

	// Activate & register cloud providers
	outliner.RegisterProvider(
		util.Validater,
		//digitalocean.Activator{},
		linode.Activator{},
		//vultr.Activator{},
	)

	err = outliner.CheckAvalible()
	if err != nil {
		panic(err)
	}
}
