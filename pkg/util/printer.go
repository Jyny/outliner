package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/logrusorgru/aurora"
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

func PrintProvidersTable(in [][]string) {
	var data [][]string
	for _, p := range in {
		data = append(data, []string{
			p[0],
			p[1],
		})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Provider",
		"API Token",
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

func PrintAPICertTable(in ol.APICert) {
	var data [][]string
	data = append(data, []string{
		in.APIurl,
		in.CertSha256,
	})
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"APIurl",
		"CertSha256",
	})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}

func PrintAPICertJSON(in ol.APICert) {
	b, err := json.Marshal(in)
	if err != nil {
		log.Println(err)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("CONGRATULATIONS! Your Outline server is up and running.")
	fmt.Println()
	fmt.Println("To manage your Outline server, please copy the following line")
	fmt.Println("(including curly brackets) into Step 2 of the Outline Manager interface:")
	fmt.Println()
	fmt.Println(aurora.Green(string(b)))
	fmt.Println()
}

func PrintCreateInstanceStart() {
	log.Println("[Creating Server]")
}

func PrintCreateInstanceWait() {
	log.Println("[Wait Server to startup]")
}

func PrintCreateInstanceDone() {
	log.Println("[Server Created]")
}

func PrintDestroyInstanceStart() {
	log.Println("[Destroying Server]")
}

func PrintDestroyInstanceDone() {
	log.Println("[Server Destroid]")
}

func PrintDeployInstanceStart() {
	log.Println("[Deploying Server]")
}

func PrintDeployInstanceWait() {
	log.Println("[Wait deployment to completed]")
}

func PrintDeployInstanceDone() {
	log.Println("[Server Deployed]")
}

func PrintVersion(v string) {
	fmt.Println(v)
}
