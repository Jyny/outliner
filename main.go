package main

import (
	"github.com/jyny/outliner/pkg/cmd"
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	cmd.Execute()
}
