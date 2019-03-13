// +build generate

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(http.Dir("./patterns"), vfsgen.Options{
		PackageName:  "assets",
		BuildTags:    "build",
		VariableName: "Patterns",
		Filename:     "assets/assets_build.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
