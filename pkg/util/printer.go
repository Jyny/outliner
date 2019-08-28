package util

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"

	ol "github.com/jyny/outliner/pkg/outliner"
)

func PrintRegionsTable(in map[string][]ol.Region) {
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

func PrintSpecsTable(in map[string][]ol.Spec) {
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

func PrintProvidersTable(in []string) {
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

func PrintInstancesTable(in []ol.Instance) {
	var data [][]string
	for _, i := range in {
		data = append(data, []string{
			i.Provider,
			i.ID,
			i.IPv4,
			i.Region.ID,
			i.Spec.ID,
		})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Provider",
		"ID",
		"IP",
		"Regien",
		"Spec",
	})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}

func PrintCreateInstanceStart() {
	fmt.Println("[Creating Instance]")
}

func PrintCreateInstanceDone() {
	fmt.Println("[Instance Created]")
}

func PrintDestroyInstanceStart(i string) {
	fmt.Println("[Searching Instance]", i)
}

func PrintDestroyInstanceDone(i string) {
	fmt.Println("[Instance Destroid]", i)
}
