package main

import (
	"github.com/shurcooL/vfsgen"
	"net/http"

	"github.com/jyny/outliner/pkg/agent/consts"
)

func main() {
	fs := http.Dir(consts.DeployScriptPath)
	opt := vfsgen.Options{
		PackageName:  consts.PackageName,
		VariableName: consts.VariableName,
	}
	err := vfsgen.Generate(fs, opt)
	if err != nil {
		panic(err)
	}
}
