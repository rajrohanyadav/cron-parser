package main 

import (
	"testing"
)

func TestPrintInterval(t *testing.T) {
	tests := []struct{
		testName string
		start int
		end int
		interval int
		expected string
	}{
		{
			testName: "Test: 1-5",
			start: 1,
			end: 5,
			interval: 1,
			expected: "1 2 3 4 5",
		},
		{
			testName: "Test: 0-10",
			start: 0,
			end: 10,
			interval: 1,
			expected: "0 1 2 3 4 5 6 7 8 9 10",
		},
		{
			testName: "Test: 1,10,2",
			start: 1,
			end: 10,
			interval: 2,
			expected: "1 3 5 7 9",
		},
		{
			testName: "Test: 10",
			start: 10,
			end: 12,
			interval: 5,
			expected: "10",
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(tt *testing.T) {
			actual := PrintIntervals(test.start, test.end, test.interval)
			if actual != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, actual)
			}
		})
	}
}
