package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("/* Reading in zip-coordinates.csv ... */")

	//	Look for the file zip-coordinates.csv and open it
	file, err := os.Open("zip-coordinates.csv")
	if err != nil {
		fmt.Println("zip-coordinates.csv not found!")
		os.Exit(0)
	}
	defer file.Close()

	//	Read in the entire csv file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("/************************/\n\n\n")

	maxLineCounter := 0

	//	For each line in the CSV file
	for i, record := range records {

		//	If we're at the beginning of a section, include the insert statement:
		if maxLineCounter == 0 {
			fmt.Println("insert into zip2location(zipcode, lat, long) values")
		}

		//	Keep track of how many lines we have processed
		maxLineCounter++

		//	Parse
		zip := record[0]
		lat := strings.TrimSpace(record[1])
		long := strings.TrimSpace(record[2])

		if i == len(records)-1 || maxLineCounter == 10000 {
			//	Use a semicolon - we're at the end
			fmt.Printf("(%v, %v, %v);\n\n", zip, lat, long)
			maxLineCounter = 0 // reset the counter
		} else {
			//	Use a comma - there is more coming
			fmt.Printf("(%v, %v, %v),\n", zip, lat, long)
		}
	}
}
