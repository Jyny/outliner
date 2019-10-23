package command

import (
	"errors"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	ol "github.com/jyny/outliner/pkg/outliner"

	// deployer agnet
	"github.com/jyny/outliner/pkg/deployer/ssh"

	// cloud provider
	"github.com/jyny/outliner/pkg/cloud/linode"
	//"github.com/jyny/outliner/pkg/cloud/vultr"
	//"github.com/jyny/outliner/pkg/cloud/digitalocean"
)

// Persistent Flags
var cfgFile string
var version = ""

// Persistent for commends
var cloud = ol.NewCloud()
var deployer = ol.NewDeployer()

// RootCmd commands
var RootCmd = &cobra.Command{
	Use:   "outliner",
	Short: "CLI tool for auto setup Outline VPN server",
	Long:  "CLI tool for auto setup Outline VPN server",
}

// Execute entry of commandline
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(initOutliner)
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "F", "", "config file (default is $HOME/.outliner/.env)")

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
	// register deployer agent
	deployer.RegisterAgent(
		ssh.NewAgent(),
	)

	// Activate & register cloud providers
	err := cloud.RegisterProvider(
		validater,
		linode.Activator{},
		//vultr.Activator{},
		//digitalocean.Activator{},
	)
	if err != nil {
		panic(err)
	}
}

func validater(actvr ol.Activator) (ol.Provider, error) {
	for _, tokenName := range actvr.ListTokenName() {
		token := viper.GetString(tokenName)
		if actvr.VerifyToken(token) {
			return actvr.GenProvider(token), nil
		}
	}
	return nil, errors.New("invalid tokens")
}
