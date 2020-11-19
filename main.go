package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// get data from sample csv
	lines, err := readCSV("sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

}

// readCSV opens a file and returns the lines
func readCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer f.Close()

	// read file
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal("Error reading lines in file", err)
	}
	return lines, nil
}
