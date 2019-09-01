package main

import (
	"github.com/jyny/outliner/pkg/runner"
	"github.com/shurcooL/vfsgen"
	"log"
	"net/http"
)

func main() {
	fs := http.Dir(runner.DeployScriptPath)
	opt := vfsgen.Options{
		PackageName:  runner.PackageName,
		VariableName: runner.VariableName,
	}
	err := vfsgen.Generate(fs, opt)
	if err != nil {
		log.Fatal(err)
	}
}
