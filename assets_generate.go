// +build generate

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(http.Dir("assets/patterns"), vfsgen.Options{
		PackageName:  "assets",
		BuildTags:    "build",
		VariableName: "Patterns",
		Filename:     "assets/patterns_build.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
