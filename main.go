package main

import (
	"bufio"
	"encoding/csv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type StockData struct {
	ID int `json:"id" gorm:"primary_key"`
	Date time.Time `json:"date"`
	Salon string `json:"salon"`
	Product string `json:"product"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
	Supplier string `json:"supplier"`
	Brand string `json:"brand"`
	Category string `json:"category"`
	SubBrand string `json:"sub_brand"`
	Type string `json:"type"`
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

func dateFormatYear(d string) (f string) {
	s := strings.Split(d, "-")
	f = "20" + s[2] + "-" + s[1] + "-" + s[0]
	return
}

func main() {
	db := dbConn()
	db.LogMode(true)
	db.DropTable(&StockData{})
	db.AutoMigrate(&StockData{})
	db.Close()

	loadProfessional()
	loadRetail()
}

func loadProfessional() {
	var files []string
	var stockTransfers []StockData
	var err error

	root := "data/stock-out/Professional"

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
		fname := filepath.Base(fileName)
		split := strings.Split(fname, " ")
		salon := split[0]
		date := "01-" + strings.Split(split[1], ".")[0]

		csvFile, _ := os.Open(fileName)

		sod := csv.NewReader(bufio.NewReader(csvFile))

		for {
			sor, err := sod.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
			}

			sl :="data/stock-list/professional.csv"

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
					formattedDate, _ := time.Parse("2006-01-02", dateFormatYear(date))
					quantity, _ := strconv.Atoi(sor[1])
					totalPrice, _ := strconv.ParseFloat(sor[2], 8)
					price := totalPrice / float64(quantity)

					stockTransfers = append(stockTransfers, StockData{Date: formattedDate, Salon: salon, Product: sor[0], Quantity: quantity, Price: price, Supplier: slr[1], Brand: slr[2], Category: slr[3], SubBrand: slr[4], Type: slr[5]})
				}
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

func loadRetail() {
	var files []string
	var stockTransfers []StockData
	var err error

	root := "data/stock-out/Retail"

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
		fname := filepath.Base(fileName)
		split := strings.Split(fname, " ")
		salon := split[0]
		date := "01-" + strings.Split(split[1], ".")[0]

		csvFile, _ := os.Open(fileName)

		sod := csv.NewReader(bufio.NewReader(csvFile))

		for {
			sor, err := sod.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
			}

			sl :="data/stock-list/retail.csv"

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

					formattedDate, _ := time.Parse("2006-01-02", dateFormatYear(date))
					quantity, _ := strconv.Atoi(sor[4])
					price, _ := strconv.ParseFloat(sor[1], 8)

					stockTransfers = append(stockTransfers, StockData{Date: formattedDate, Salon: salon, Product: sor[0], Quantity: quantity, Price: price, Supplier: slr[1], Brand: slr[2], Category: slr[3], SubBrand: slr[4], Type: slr[5]})
				}
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