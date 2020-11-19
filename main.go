package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// get data from sample csv
	lines, err := readCSV("sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	writeCSV(lines)

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

// writeCSV iterates through lines from parsed file and sends to new file, currently as-is
func writeCSV(lines [][]string) {
	file, err := os.Create("clean-sample.csv")
	if err != nil {
		log.Fatal("Failed to create file: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, line := range lines {
		_ = writer.Write(line)
	}
}
