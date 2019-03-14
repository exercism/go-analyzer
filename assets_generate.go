// +build generate

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(http.Dir("patterns"), vfsgen.Options{
		PackageName:  "analyzer",
		BuildTags:    "build",
		VariableName: "Patterns",
		Filename:     "analyzer/patterns_build.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
