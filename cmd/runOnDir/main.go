package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/exercism/go-analyzer/analyzer"
)

var (
	exercise  = flag.String("exercise", "", "exercise slug (e.g. 'two-fer'). All sub-dirs must contain this exercise!")
	parentDir = flag.String("parentDir", "", "run analyzer for all sub-directories in this folder")
	output    = flag.String("output", "analysis.json", "name of the output file")
)

func main() {
	flag.Parse()
	if *exercise == "" || *parentDir == "" || *output == "" {
		log.Println("invalid empty parameter")
		flag.Usage()
		return
	}

	dirs, err := analyzer.GetDirs(".", http.Dir(*parentDir))
	if err != nil {
		log.Fatal(err)
	}

	sum := map[analyzer.Status]int{}
	for _, dir := range dirs {
		res := analyzer.Analyze(*exercise, path.Join(*parentDir, dir))
		for _, err := range res.Errors {
			log.Println(err)
		}
		if len(res.Errors) != 0 {
			continue
		}

		sum[res.Status]++
		bytes, err := json.MarshalIndent(res, "", "\t")
		if err != nil {
			log.Println(err)
			continue
		}
		if err := ioutil.WriteFile(path.Join(*parentDir, dir, *output), append(bytes, '\n'), 0644); err != nil {
			log.Println(err)
			continue
		}
	}

	fmt.Println("Statistics:")
	fmt.Printf("%+v\n", sum)
}
