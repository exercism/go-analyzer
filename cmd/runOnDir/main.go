package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/exercism/go-analyzer/analyzer"
	"github.com/exercism/go-analyzer/assets"
)

var (
	exercise     = flag.String("exercise", "", "exercise slug (e.g. 'two-fer'). All sub-dirs must contain this exercise!")
	parentDir    = flag.String("parentDir", "", "run analyzer for all sub-directories in this folder")
	output       = flag.String("output", "analysis.json", "name of the output file")
	printStatus  = flag.String("printStatus", "", "prints out every folder with the provided status")
	minPrintDiff = flag.Float64("minPrintDiff", 1, "prints out all diffs above set rating")
)

func main() {
	flag.Parse()
	if *exercise == "" || *parentDir == "" || *output == "" {
		log.Println("invalid empty parameter")
		flag.Usage()
		return
	}

	dirs, err := assets.GetDirs(".", http.Dir(*parentDir))
	if err != nil {
		log.Fatal(err)
	}
	fileName := strings.ReplaceAll(*exercise, "-", "_") + ".go:1"

	sum := map[analyzer.Status]int{}
	for _, dir := range dirs {
		res := analyzer.Analyze(*exercise, path.Join(*parentDir, dir))
		for _, err := range res.Errors {
			log.Printf("ERROR on %s:\n", path.Join(*parentDir, dir))
			log.Println(err)
		}

		if res.Status == analyzer.Status(*printStatus) {
			var minDiff string
			if *minPrintDiff <= res.Rating {
				minDiff = res.MinDiff
			}
			fmt.Printf("Status %s: severity: %d, rating: %.2f, path: %s (%s)\n%s",
				*printStatus,
				res.Severity,
				res.Rating,
				path.Join(*parentDir, dir, fileName),
				path.Join(*parentDir, dir, *output+":1"),
				minDiff,
			)
		}

		bytes, err := json.MarshalIndent(res, "", "\t")
		if err != nil {
			log.Println(err)
			eject(sum, dir)
			continue
		}
		if err := ioutil.WriteFile(path.Join(*parentDir, dir, *output), append(bytes, '\n'), 0644); err != nil {
			log.Println(err)
			eject(sum, dir)
			continue
		}
		sum[res.Status]++
	}

	fmt.Println("Statistics:")
	fmt.Printf("%+v\n", sum)
}

func eject(sum map[analyzer.Status]int, dir string) {
	sum[analyzer.Ejected]++
	if *printStatus == "ejected" {
		fmt.Printf("Status %s: %s\n", *printStatus, path.Join(*parentDir, dir))
	}
}
