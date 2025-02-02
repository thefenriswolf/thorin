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

// Jahr;Monat;Tag;Kategorie;Betrag;Notiz;Typ
const (
	yearPos = iota
	monthPos
	dayPos
	catPos
	amountPos
	notePos
	directionPos
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
	data := make([]entry, 0, fle)
	for lineNum, lineContent := range fc {
		if lineNum > 0 { // skip first line
			data = append(data, parseLine(lineContent))
		}
	}
	sort.SliceStable(data, func(i, j int) bool { return data[i].date.Before(data[j].date) })

	var sum float64
	for i := range data {
		sum = sum + data[i].amount
		fmt.Println(
			data[i].date.Format(outputDateFormat),
			data[i].category,
			data[i].amount,
			data[i].note,
			data[i].income)
	}
	fmt.Println("-----------------------------------------")
	fmt.Printf("Sum: %.2f\n", sum)
}
