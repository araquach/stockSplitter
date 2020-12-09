package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
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

func split() {
	//file := "stock.csv"
	//
	//data, err := ioutil.ReadFile(file)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//section := strings.Split(string(data), "Product Name")
	//
	//var combined []string
	//
	//for _, split := range section {
	//	p := strings.Split(split, "\n")
	//
	//	if len(p) > 4 {
	//		combined = append(combined, split)
	//		s, _ := strconv.Atoi(f[0])
	//
	//		info := strings.Split(p[0], ",")
	//		ratingLine := strings.Split(p[2], ",")
	//		rating := strings.Split(ratingLine[0], " ")
	//		date, _ := time.Parse("2006-01-02", dateFormatYear(info[0]))
	//		stylist := info[2]
	//		client := info[1]
	//		review := strings.Trim(strings.Split(p[3], ",")[0], "\"")
	//		ratingInt, _ := strconv.Atoi(rating[1])
	//
	//		if len(review) > 25 && ratingInt >= 4 {
	//			{
	//				reviews = append(reviews, ClientReview{Date: date, Review: review, Client: client, Stylist: stylist, Salon: s})
	//			}
	//		}
	//	}
	//}
}
