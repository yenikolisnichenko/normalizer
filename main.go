package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"
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

		timeStamp := changeTime(line[0])
		address := line[1]
		comment := line[7]

		normLine := []string{
			timeStamp,
			address,
			comment,
		}

		_ = writer.Write(normLine)
	}
}

// changeTime takes in a timestamp and converts it to RFC3339 and EST timezone
func changeTime(timestamp string) string {

	const format = "01/02/06 03:04:05 PM"

	// get current timezone
	PST, _ := time.LoadLocation("America/Los_Angeles")

	// get desired timezone
	EST, _ := time.LoadLocation("America/New_York")

	t, err := time.ParseInLocation(format, timestamp, PST)
	if err != nil {
		log.Fatal("Error parsing time in location: ", err)
	}
	// get timestamp in RFC3339 format and EST timezone
	t.Format(time.RFC3339)
	estTime := t.In(EST)
	return estTime.String()
}
