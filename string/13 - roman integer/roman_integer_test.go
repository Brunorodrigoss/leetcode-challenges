package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	var tests = []struct {
		symbol string
		want   int
	}{
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},

		{"III", 3},
		{"IV", 4},
		{"XXII", 22},
		{"XXVII", 27},

		{"LX", 60},
		{"LVIII", 58},
		{"LXXIV", 74},
		{"MCMXCIV", 1994},
	}

	for _, test := range tests {
		t.Run(test.symbol, func(t *testing.T) {
			result := romanToInt(test.symbol)

			if result != test.want {
				t.Errorf("romanToInt(\"%s\") = %d; want %d", test.symbol, result, test.want)
			}
		})
	}
}
