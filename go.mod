module github.com/jyny/outliner

go 1.12

require (
	github.com/Songmu/prompter v0.2.0
	github.com/digitalocean/godo v1.19.0
	github.com/linode/linodego v0.10.0
	github.com/logrusorgru/aurora v0.0.0-20190803045625-94edacc10f9b
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/olekukonko/tablewriter v0.0.1
	github.com/pkg/sftp v1.10.1
	github.com/sethvargo/go-password v0.1.2
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/vultr/govultr v0.1.4
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/text v0.3.8 // indirect
)

replace github.com/spf13/viper v1.4.0 => github.com/jyny/viper v1.4.1
