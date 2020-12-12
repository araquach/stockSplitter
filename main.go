package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type StockData struct {
	ID int `json:"id" gorm:"primary_key"`
	Product string `json:"product"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	Supplier string `json:"supplier"`
	Brand string `json:"brand"`
	Category string `json:"category"`
	SubBrand string `json:"sub_brand"`
	Retail string `json:"retail"`
}

func dbConn() (db *gorm.DB) {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	return db
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db := dbConn()
	db.LogMode(true)
	db.AutoMigrate(&StockData{})
	db.Close()

	//getStockOut()
	//getStockList()

	combine()
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
	var stockTransfers []StockData
	var err error

	so := "data/stock-out.csv"

	csvFile, _ := os.Open(so)

	sod := csv.NewReader(bufio.NewReader(csvFile))

	for {
		sor, err := sod.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		if len(sor[1]) > 0 && !strings.Contains(sor[0], "Page") &&
			!strings.Contains(sor[1], "#") &&
			!strings.Contains(sor[0], "Total"){
		}

		sl :="data/stock-list.csv"

		csvFile2, _ := os.Open(sl)

		sld := csv.NewReader(bufio.NewReader(csvFile2))

		for {
			slr, err := sld.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
			}

			if slr[0] == sor[0] {

				quantity, _ := strconv.Atoi(sor[1])
				price, _ := strconv.Atoi(sor[2])

				stockTransfers = append(stockTransfers, StockData{Product: sor[0], Quantity: quantity, Price: price, Supplier: slr[1], Brand: slr[2], Category: slr[3], SubBrand: slr[4], Retail: slr[5]})

			}
		}
	}
	for _, r := range stockTransfers {
		db := dbConn()
		db.LogMode(true)
		db.Create(&r)
		if err != nil {
			log.Println(err)
		}
		db.Close()
	}
}