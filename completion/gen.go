package main

import (
	"github.com/jyny/outliner/command"
)

func main() {
	err := command.RootCmd.GenBashCompletionFile("build/outliner_bash_completion")
	if err != nil {
		panic(err)
	}
	err = command.RootCmd.GenZshCompletionFile("build/outliner_zsh_completion")
	if err != nil {
		panic(err)
	}
}
