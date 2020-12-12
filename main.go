package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	getStockOut()
	getStockList()
}

func getStockOut() {
	file := "data/stock-out.csv"

	csvFile, _ := os.Open(file)

	r := csv.NewReader(bufio.NewReader(csvFile))

	fmt.Println("-- Stock Out --")

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		product := record[0]
		quantity := record[1]

		if len(record[1]) > 0 && !strings.Contains(record[0], "Page") &&
			!strings.Contains(record[1], "#") &&
			!strings.Contains(record[0], "Total"){
			fmt.Printf("%v - Quantity Used: %v\n", product, quantity)
		}
	}
}

func getStockList() {
	file := "data/stock-list.csv"

	csvFile, _ := os.Open(file)

	r := csv.NewReader(bufio.NewReader(csvFile))

	fmt.Println("-- Stock List --")

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		product := record[0]
		supplier := record[1]
		brand := record[2]
		cat := record[3]
		sub := record[4]
		retail := record[5]
		if retail == "1" {
			retail = "Yes"
		} else {
			retail = "No"
		}

		fmt.Printf("%v - %v - %v - %v - %v - %v\n", product, supplier, brand, cat, sub, retail)
	}
}

func combine() {

}