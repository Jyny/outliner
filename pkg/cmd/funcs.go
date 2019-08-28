package cmd

import (
	"errors"

	"github.com/spf13/viper"

	ol "github.com/jyny/outliner/pkg/outliner"
)

func validater(actvr ol.Activator) (ol.Provider, error) {
	for _, tokenName := range actvr.ListTokenName() {
		viper.SetDefault(tokenName, "")
		token := viper.Get(tokenName).(string)
		if actvr.VerifyToken(token) {
			return actvr.GenProvider(), nil
		}
	}
	return nil, errors.New("invalid tokens")
}
