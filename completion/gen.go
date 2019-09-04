package main

import (
	"github.com/jyny/outliner/pkg/cmd"
)

func main() {
	err := cmd.RootCmd.GenBashCompletionFile("build/outliner_bash_completion")
	if err != nil {
		panic(err)
	}
	err = cmd.RootCmd.GenZshCompletionFile("build/outliner_zsh_completion")
	if err != nil {
		panic(err)
	}
}
