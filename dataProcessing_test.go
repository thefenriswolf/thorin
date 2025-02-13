package main

import (
	"testing"
	"time"
)

func TestIsIncome(t *testing.T) {
	var tests = []struct {
		a    string
		want bool
	}{
		{"Ausgabe", false},
		{"Einnahme", true},
		{"dafasdfa", false},
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

func TestParseDate(t *testing.T) {
	var tests = []struct {
		d    string
		m    string
		y    string
		want time.Time
	}{
		{"01", "02", "2024", time.Date(2024, 02, 01, 0, 0, 0, 0, time.UTC)},
		{"31", "01", "2024", time.Date(2024, 01, 31, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.d+"."+tt.m+"."+tt.y, func(t *testing.T) {
			ans := parseDate(tt.y, tt.m, tt.d)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
