package main

import (
	"testing"
)

func TestIsIncome(t *testing.T) {
	var tests = []struct {
		a    string
		want bool
	}{
		{"Ausgabe", false},
		{"Einnahme", true},
	}
	for _, tt := range tests {
		testname := tt.a
		t.Run(testname, func(t *testing.T) {
			ans := isIncome(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

/* // FIX: time field must be null
func TestParseLine(t *testing.T) {
	var date time.Time = time.Date(2024, 2, 15, 0, 0, 0, 0, time.FixedZone("UTC-8", -8*60*60))
	var lineResult = entry{
		date:     date,
		category: "TestCategory",
		amount:   314.41,
		note:     "TestNote",
		income:   true,
	}
	var testLine = []string{"2024", "2", "15", "TestCategory", "314.41", "TestNote", "true"}
	ans := parseLine(testLine)
	if ans != lineResult {
		t.Errorf("parseLine = %v; want %v", ans, lineResult)
	}

}
*/
/*
func TestgetFileContents(t *testing.T) //fileName string) ([][]string, error) {

func TestgetDirection(t *testing.T) //income string) bool {

func TestparseDate(t *testing.T) //y, m, d string) time.Time {

*/
