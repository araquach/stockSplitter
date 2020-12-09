package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	split2()
}

func split1() {
	file := "stock.csv"

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	section := strings.SplitAfter(string(data), "Brand ")

	fmt.Println(section[4])

	// var combined []string

	for _, split := range section[2:] {
		p := strings.Split(split, "\n")[0]
		lineSplit := strings.Split(p, " -")

		supplier := lineSplit[0]

		cat := strings.TrimLeft(lineSplit[1], "Category ")
		cat = strings.TrimRight(cat, ",")

		fmt.Printf("Supplier: %v\nCategory: %v\n\n", supplier, cat)
	}
}

func split2() {
	var err error
	var files []string

	root := "data"

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".csv" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	for _, fileName := range files {
		fname := strings.Split(fileName, "/")[1]

		fileBytes, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Println(err)
		}

		split := strings.SplitAfter(string(fileBytes), "\n")

		var section []string

		for _, line := range split {
			if !strings.Contains(line, "Page") {
				section = append(section, line)
			}
		}

		joined := strings.Join(section, "")

		for _, l := range joined {

			s := strings.SplitAfter(string(l), "\n")

			csvReady := strings.Join(s, "")

			r := csv.NewReader(strings.NewReader(csvReady))

			for {
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Println(err)
				}
				fmt.Println(record)
			}
		}
	}
}
