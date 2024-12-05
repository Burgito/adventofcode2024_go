package main

import (
	"adventofcode/arrays"
	"adventofcode/inputreaders"
	"fmt"
	"strconv"
)

func iAbsDiff(a int, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

func isAllIncrease(slice []int) bool {
	length := len(slice)
	for i := 0; i < length; i++ {
		if i < length-1 {
			if slice[i] >= slice[i+1] {
				return false
			}
		}
	}

	return true
}

func isSafeByVacuity(slice []int) bool {
	length := len(slice)
	if length == 1 || length == 0 {
		return true
	}

	return false
}

func levelsAreUnsafe(increase bool, value1 int, value2 int) bool {
	if value1 == value2 {
		return true
	}

	if increase && (value1 > value2) {
		return true
	}

	if !increase && (value1 < value2) {
		return true
	}

	if iAbsDiff(value1, value2) > 3 {
		return true
	}

	return false
}

func isReportSafe(slice []int) bool {
	length := len(slice)

	if isSafeByVacuity(slice) {
		return true
	}

	increase := false
	if slice[0] < slice[1] {
		increase = true
	}

	for i := 0; i < length-1; i++ {
		if levelsAreUnsafe(increase, slice[i], slice[i+1]) {
			return false
		}
	}

	return true
}

func part1() {
	reports := inputreaders.ReadIntegersPerLines("./input.txt", " ")
	totalSafe := 0
	for _, report := range reports {
		if isReportSafe(report) {
			totalSafe += 1
		}
	}

	fmt.Println("Total safe reports : " + strconv.Itoa(totalSafe))
}

func part2() {
	reports := inputreaders.ReadIntegersPerLines("./input.txt", " ")
	totalSafe := 0
	for _, report := range reports {
		if isReportSafe(report) {
			totalSafe += 1
		} else {
			for idx := range report {
				if isReportSafe(arrays.RemoveIndexFromIntArray(report, idx)) {
					totalSafe += 1
					break
				}
			}
		}
	}

	fmt.Println("Total safe reports : " + strconv.Itoa(totalSafe))

}

func main() {
	part1()
	part2()
}
