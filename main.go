package main

import (
	"github.com/jyny/outliner/command"
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	command.Execute()
}
