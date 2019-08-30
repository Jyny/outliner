package cmd

import (
	"errors"
	"os"

	"github.com/olekukonko/tablewriter"

	"github.com/spf13/viper"

	ol "github.com/jyny/outliner/pkg/outliner"
)

func validater(actvr ol.Activator) (ol.Provider, error) {
	for _, tokenName := range actvr.ListTokenName() {
		token := viper.GetString(tokenName)
		if actvr.VerifyToken(token) {
			return actvr.GenProvider(token), nil
		}
	}
	return nil, errors.New("invalid tokens")
}

func printRegions(in map[string][]ol.Region) {
	var data [][]string
	for pname, p := range in {
		for _, r := range p {
			data = append(data, []string{
				pname,
				r.ID,
				r.Note,
			})
		}
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Provider",
		"Region",
		"Note",
	})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}

func printSpecs(in map[string][]ol.Spec) {
	var data [][]string
	for pname, p := range in {
		for _, s := range p {
			data = append(data, []string{
				pname,
				s.ID,
				s.Transfer,
				s.Price,
			})
		}
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Provider",
		"SpecID",
		"Transfer",
		"Price",
	})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}

func printProvider(in []string) {
	var data [][]string
	for _, p := range in {
		data = append(data, []string{
			p,
		})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Provider",
	})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}
