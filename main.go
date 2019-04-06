//go:generate go run --tags=generate assets_generate.go

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/exercism/go-analyzer/analyzer"
	"github.com/namsral/flag"
)

var (
	exercise     = flag.String("exercise", "", "exercise slug (e.g. 'two-fer')")
	solutionPath = flag.String("solution", "", "path to solution to be processed")
	output       = flag.String("output", "analysis.json", "name of the output file")
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
	log.Printf("Starting check for `%s` exercise in folder `%s`\n", *exercise, *solutionPath)

	res := analyzer.Analyze(*exercise, *solutionPath)
	for _, errStr := range res.Errors {
		log.Println(errStr)
	}
	bytes, err := toJSON(res)
	if err != nil {
		log.Printf("%+v", err)
		os.Exit(2)
	}

	outputFile := path.Join(*solutionPath, *output)
	if err := ioutil.WriteFile(outputFile, append(bytes, '\n'), 0644); err != nil {
		log.Printf("%+v", err)
		os.Exit(3)
	}
	log.Printf("Output written to %s", outputFile)
}

func toJSON(res interface{}) ([]byte, error) {
	return json.MarshalIndent(res, "", "\t")
}
