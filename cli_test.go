package main

import "testing"

func TestReadTimeframe(t *testing.T) {
	var tests = []struct {
		tf   string
		want [6]int
	}{
		{"01.01.2024:30.01.2024", [6]int{2024, 1, 1, 2024, 1, 30}},
		{"1.1.2024:30.1.2024", [6]int{2024, 1, 1, 2024, 1, 30}},
		{"00.00.2024:00.00.2025", [6]int{2024, 0, 0, 2025, 0, 0}},
		{"0.0.2024:0.0.2025", [6]int{2024, 0, 0, 2025, 0, 0}},
		{"0.01.2024:0.02.2025", [6]int{2024, 1, 0, 2025, 2, 0}},
		{"0.1.2024:0.2.2025", [6]int{2024, 1, 0, 2025, 2, 0}},
	}
	for _, tt := range tests {
		testname := tt.tf
		t.Run(testname, func(t *testing.T) {
			ans := readTimeframe(tt.tf)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}

/*
func TestParseArgs(t *testing.T) {
	var args = []string{
		"-h",
		"--help",
		"-f", "test.csv",
		"--file", "test.csv",
		"-t", "01.01.2024:30.01.2024",
	}
	for _, tt := range args[:1] {
		t.Run(tt, func(t *testing.T) {
			ans := parseArgs(tt)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
*/
