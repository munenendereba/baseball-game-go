package main

import (
	"fmt"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	var tests = []struct {
		ops []string
		sum int
	}{
		{[]string{"5", "2", "C", "D", "+"}, 30},
		{[]string{"1"}, 1},
		{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
	}

	for i, test := range tests {
		ops := test.ops
		sum := test.sum

		testname := fmt.Sprintf("running test %d", i)

		t.Run(testname, func(t *testing.T) {
			res := CalculatePoints(ops)

			if res != sum {
				t.Errorf(`Calculate Points(%s), expected result to be %d, got %d`, ops, sum, res)
			}
		})

	}
}
