package main

import (
	"github.com/jyny/outliner/pkg/agent"
	"github.com/shurcooL/vfsgen"
	"log"
	"net/http"
)

func main() {
	fs := http.Dir(agent.DeployScriptPath)
	opt := vfsgen.Options{
		PackageName:  agent.PackageName,
		VariableName: agent.VariableName,
	}
	err := vfsgen.Generate(fs, opt)
	if err != nil {
		log.Fatal(err)
	}
}
