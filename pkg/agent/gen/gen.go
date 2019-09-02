package main

import (
	"github.com/jyny/outliner/pkg/agent/constdef"
	"github.com/shurcooL/vfsgen"
	"log"
	"net/http"
)

func main() {
	fs := http.Dir(constdef.DeployScriptPath)
	opt := vfsgen.Options{
		PackageName:  constdef.PackageName,
		VariableName: constdef.VariableName,
	}
	err := vfsgen.Generate(fs, opt)
	if err != nil {
		log.Fatal(err)
	}
}
