package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	yearPos = iota
	monthPos
	dayPos
	catPos
	amountPos
	notePos
	directionPos
)
const (
	yearStart = iota
	monthStart
	dayStart
	yearEnd
	monthEnd
	dayEnd
)

const inputDateFormat = "2006 1 2"
const outputDateFormat = "02.01.2006"

type entry struct {
	date     time.Time
	category string
	amount   float64
	note     string
	income   bool
}

func getFileContents(fileName string) ([][]string, error) {
	fileHandle, ferr := os.Open(fileName)
	defer fileHandle.Close()
	if ferr != nil {
		err := fmt.Errorf("Could not open file: %v", ferr)
		return nil, err
	}
	r := csv.NewReader(fileHandle)
	r.Comma = ';'
	//r.Comment='#'
	r.FieldsPerRecord = 7
	r.TrimLeadingSpace = true
	fc, rerr := r.ReadAll()
	if rerr != nil {
		err := fmt.Errorf("Could not read file contents: %v", rerr)
		return nil, err
	}
	return fc, nil
}
func isIncome(direction string) bool {
	var income bool
	if direction == "Ausgabe" {
		income = false
	}
	if direction == "Einnahme" {
		income = true
	}
	return income
}

func parseDate(y, m, d string) time.Time {
	dateString := y + " " + m + " " + d
	date, err := time.Parse(inputDateFormat, dateString)
	if err != nil {
		log.Fatal("Could not parse date: ", err)
	}
	return date
}

func parseLine(line []string) entry {
	date := parseDate(line[yearPos], line[monthPos], line[dayPos])
	//fmt.Println("date: ", date.Format(outputDateFormat))
	category := line[catPos]
	amount, _ := strconv.ParseFloat(strings.Replace(line[amountPos], ",", ".", 1), 64)
	note := line[notePos]
	income := isIncome(line[directionPos])
	if !income {
		amount = amount * -1
	}

	var lineData = entry{date: date,
		category: category,
		amount:   amount,
		note:     note,
		income:   income}
	return lineData
}

func collectData(fc [][]string, dlen int) []entry {
	data := make([]entry, 0, dlen)
	for lineNum, lineContent := range fc {
		if lineNum > 0 { // skip first line
			data = append(data, parseLine(lineContent))
		}
	}
	sort.SliceStable(data, func(i, j int) bool { return data[i].date.Before(data[j].date) })
	return data
}

func genSum(data []entry, start int, end int) float64 {
	var sum float64
	if end > len(data)-1 {
		end = len(data) - 1
	}
	for i := start; i <= end; i++ {
		sum = sum + data[i].amount
	}
	return sum
}

func getTimeframe(data []entry, timeframe [6]int) ([]entry, error) {
	// TODO: simplify
	// FIX: 2023 -> 2024 does not work
	var subset []entry
	// check if timeframe was specified at all
	if timeframe[yearStart] == 0 && timeframe[monthStart] == 0 && timeframe[dayStart] == 0 && timeframe[yearEnd] == 0 && timeframe[monthEnd] == 0 && timeframe[dayEnd] == 0 {
		return nil, fmt.Errorf("Invalid interval or timeframe specified")
	}
	// check if timeframe is not a interval
	if timeframe[yearEnd] == 0 && timeframe[monthEnd] == 0 && timeframe[dayEnd] == 0 {
		// get year
		if timeframe[dayStart] == 0 && timeframe[monthStart] == 0 && timeframe[yearStart] != 0 {
			for i := range data {
				y, _, _ := data[i].date.Date()
				if y == timeframe[yearStart] {
					subset = append(subset, data[i])
				}
			}
			sort.SliceStable(data, func(i, j int) bool { return data[i].date.Before(data[j].date) })
			return subset, nil
		}
		// get month
		if timeframe[dayStart] == 0 && timeframe[monthStart] != 0 && timeframe[yearStart] != 0 {
			for i := range data {
				y, m, _ := data[i].date.Date()
				if y == timeframe[yearStart] && int(m) == timeframe[monthStart] {
					subset = append(subset, data[i])
				}
			}
			sort.SliceStable(data, func(i, j int) bool { return data[i].date.Before(data[j].date) })
			return subset, nil
		}
		// get day
		if timeframe[dayStart] != 0 && timeframe[monthStart] != 0 && timeframe[yearStart] != 0 {
			for i := range data {
				y, m, d := data[i].date.Date()
				if y == timeframe[yearStart] && int(m) == timeframe[monthStart] && d == timeframe[dayStart] {
					subset = append(subset, data[i])
				}
			}
			sort.SliceStable(data, func(i, j int) bool { return data[i].date.Before(data[j].date) })
			return subset, nil
		}

	}
	if timeframe[yearEnd] != 0 || timeframe[monthEnd] != 0 || timeframe[dayEnd] != 0 {
		// get year
		if timeframe[dayStart] == 0 && timeframe[monthStart] == 0 && timeframe[yearStart] != 0 {
			for i := range data {
				y, _, _ := data[i].date.Date()
				if y >= timeframe[yearStart] && y <= timeframe[yearEnd] {
					subset = append(subset, data[i])
				}
			}
			sort.SliceStable(data, func(i, j int) bool { return data[i].date.Before(data[j].date) })
			return subset, nil
		}
		// get month
		if timeframe[dayStart] == 0 && timeframe[monthStart] != 0 && timeframe[yearStart] != 0 {
			for i := range data {
				y, m, _ := data[i].date.Date()
				if y >= timeframe[yearStart] && int(m) >= timeframe[monthStart] && y <= timeframe[yearEnd] && int(m) <= timeframe[monthEnd] {
					subset = append(subset, data[i])
				}
			}
			sort.SliceStable(data, func(i, j int) bool { return data[i].date.Before(data[j].date) })
			return subset, nil
		}
		// get day
		if timeframe[dayStart] != 0 && timeframe[monthStart] != 0 && timeframe[yearStart] != 0 {
			for i := range data {
				y, m, d := data[i].date.Date()
				if y >= timeframe[yearStart] && int(m) >= timeframe[monthStart] && d >= timeframe[dayStart] && y <= timeframe[yearEnd] && int(m) <= timeframe[monthEnd] && d <= timeframe[dayEnd] {
					subset = append(subset, data[i])
				}
			}
			sort.SliceStable(data, func(i, j int) bool { return data[i].date.Before(data[j].date) })
			return subset, nil
		}
	}
	return nil, nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("No input file specified")
	}
	fp := args[1]
	fc, err := getFileContents(fp)
	if err != nil {
		log.Fatal(err)
	}
	fle := len(fc)
	data := collectData(fc, fle)
	d, err := getTimeframe(data, [6]int{2024, 10, 1, 2024, 10, 5})
	if err != nil {
		log.Fatal(err)
	}
	for i := range d {
		fmt.Println(
			d[i].date.Format(outputDateFormat),
			d[i].category,
			d[i].amount,
			d[i].note,
			d[i].income)
	}
	fmt.Println("-----------------------------------------")
	sum := genSum(d, 0, len(data))
	fmt.Printf("Sum: %.2f\n", sum)
}
