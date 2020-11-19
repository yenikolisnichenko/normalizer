package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	for i, line := range lines {

		// pass the header as is
		if i == 0 {
			_ = writer.Write(line)
			continue
		}

		timeStamp := changeTime(line[0])
		address := line[1]
		zipcode := checkZip(line[2])
		fullname := upcaseFn(line[3])
		fooDur := durationSecs(line[4])
		barDur := durationSecs(line[5])
		totalDur := fooDur + barDur
		comment := line[7]

		normLine := []string{
			timeStamp,
			address,
			zipcode,
			fullname,
			strconv.FormatFloat(fooDur, 'f', 3, 64),
			strconv.FormatFloat(barDur, 'f', 3, 64),
			strconv.FormatFloat(totalDur, 'f', 3, 64),
			comment,
		}

		_ = writer.Write(normLine)
	}
}

// changeTime takes in a timestamp and converts it to RFC3339 and EST timezone
func changeTime(ts string) string {
	const format = "01/02/06 03:04:05 PM"

	// parse string and pad with zeros before passing it down
	timestamp := padTimestamp(ts)

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

// padTimestamp pads date and time so that changeTime() works as expected
func padTimestamp(timestamp string) string {
	// split timestamp string
	splitStr := strings.Split(timestamp, " ")

	// grab date and time to modify
	date := splitStr[0]
	time := splitStr[1]

	// split on the / to check if two digit, pad if needed
	padDate := padZeros(date, "/")
	padTime := padZeros(time, ":")

	// after checking date, join back together and replace value on main []
	splitStr[0] = padDate
	splitStr[1] = padTime

	// join main [] to return as a string
	paddedTimestamp := strings.Join(splitStr, " ")
	return paddedTimestamp
}

// padZeros checks for single digit string and add zero
func padZeros(str string, c string) string {
	splitStr := strings.Split(str, c)
	for i, dd := range splitStr {
		s := fmt.Sprintf("%02s", dd)
		splitStr[i] = s
	}
	joinStr := strings.Join(splitStr, c)

	return joinStr
}

// checkZip converts zip str to 5 char long zip
func checkZip(zip string) string {
	// pad with zeros if necessary
	paddedZip := fmt.Sprintf("%05s", zip)
	return paddedZip
}

// upcaseFn upcases fullname
func upcaseFn(fullname string) string {
	// ToUpper returns s with all Unicode letters mapped to their upper case, handles non-English names
	return strings.ToUpper(fullname)
}

// durationSecs splits duration into hour, min, sec and returns total seconds
func durationSecs(dur string) float64 {
	durSlice := strings.Split(dur, ":")

	hrs, _ := strconv.ParseFloat(durSlice[0], 64)
	mins, _ := strconv.ParseFloat(durSlice[1], 64)
	secs, _ := strconv.ParseFloat(durSlice[2], 64)

	hrMins := hrs * 60.0
	minSecs := (mins + hrMins) * 60.0
	totalSecs := secs + minSecs
	return float64(totalSecs)
}
