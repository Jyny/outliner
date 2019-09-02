package cmd

import (
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jyny/outliner/pkg/agent"
	ol "github.com/jyny/outliner/pkg/outliner"
	"github.com/jyny/outliner/pkg/util"

	//"github.com/jyny/outliner/pkg/cloud/digitalocean"
	"github.com/jyny/outliner/pkg/cloud/linode"
	//"github.com/jyny/outliner/pkg/cloud/vultr"
)

// Persistent Flags
var cfgFile string

// Persistent for commends
var cloud = ol.NewCloud()
var deployer = ol.NewDeployer()

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
	cobra.OnInitialize(initOutliner)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "F", "", "config file (default is $HOME/.outliner/.env)")

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
}

func initOutliner() {
	// add new agent to deployer
	deployer.Init(agent.New())

func initProvider() {
	// Activate & register cloud providers
	err := cloud.RegisterProvider(
		util.Validater,
		//digitalocean.Activator{},
		linode.Activator{},
		//vultr.Activator{},
	)
	if err != nil {
		panic(err)
	}
}
