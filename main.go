//go:generate go run --tags=generate assets_generate.go

package main

import (
	"log"
	"os"

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

	solution := analyzer.LoadPackage(*solutionPath)
	goodPattern, err := analyzer.CheckPattern(*exercise, solution)
	if err != nil {
		log.Println(err)
	}

	// 	TODO:
	// 	 - import and run suggest package on solution (can be 2nd step)
	// 	 - evaluate result
	// 	 - write result to file
	_ = goodPattern
}
