module github.com/jyny/outliner

go 1.12

require (
	github.com/Songmu/prompter v0.2.0
	github.com/digitalocean/godo v1.19.0
	github.com/linode/linodego v0.10.0
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/olekukonko/tablewriter v0.0.1
	github.com/sethvargo/go-password v0.1.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/vultr/govultr v0.1.4
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sys v0.0.0-20190826190057-c7b8b68b1456 // indirect
	golang.org/x/text v0.3.2 // indirect
)

replace github.com/spf13/viper v1.4.0 => github.com/jyny/viper v1.4.1
