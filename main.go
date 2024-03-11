package main

import (
	"fmt"
	"strconv"
)

func CalculatePoints(ops []string) int {
	var res int = 0

	numbers := []int{}
	removed, j := 0, 0

	for i, op := range ops {
		j = i - removed
		switch op {
		case "+":
			numbers = append(numbers, numbers[j-1]+numbers[j-2])
		case "D":
			numbers = append(numbers, numbers[j-1]*2)
		case "C":
			numbers = numbers[:j-1]
			removed += 2
		default:
			num, _ := strconv.Atoi(op)
			numbers = append(numbers, num)
		}
	}

	for _, num := range numbers {
		res += num
	}

	return res
}

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	fmt.Println(CalculatePoints(ops))
}
