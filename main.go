//go:generate go run --tags=generate assets_generate.go

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/exercism/go-analyzer/analyzer"
	"github.com/namsral/flag"
)

var (
	exercise     = flag.String("exercise", "", "exercise slug (e.g. 'two-fer')")
	solutionPath = flag.String("solution", "", "path to solution to be processed")
)

func main() {
	flag.Parse()
	if *exercise == "" || *solutionPath == "" {
		if flag.NArg() < 2 {
			flag.Usage()
			os.Exit(1)
		}
		args := flag.Args()
		*exercise, *solutionPath = args[0], args[1]
	}

	res := analyzer.Analyze(*exercise, *solutionPath)
	bytes, err := json.MarshalIndent(res, "", "\t")
	if err != nil {
		os.Exit(2)
	}
	if err := ioutil.WriteFile(path.Join(*solutionPath, "analysis.json"), bytes, 0644); err != nil {
		os.Exit(3)
	}
}
