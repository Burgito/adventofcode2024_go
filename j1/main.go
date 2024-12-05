package main

import (
	"adventofcode/inputreaders"
	"adventofcode/ints"
	"fmt"
	"slices"
)

func getTotalDistance(sliceLeft []int, sliceRight []int) int {
	var total int
	for idx, value := range sliceLeft {
		total += ints.IAbsDiff(value, sliceRight[idx])
	}

	return total
}

func getTotalNumberOfOccurences(sliceLeft []int, sliceRight []int) int {
	var total int
	for _, valueL := range sliceLeft {
		for _, valueR := range sliceRight {
			if valueR == valueL {
				total += valueL
			}
		}
	}

	return total
}

func part1() {
	lefts := []int{}
	rights := []int{}
	lines := inputreaders.ReadIntegersPerLines("./input.txt", "   ")
	for _, line := range lines {
		lefts = append(lefts, line[0])
		rights = append(rights, line[1])
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	var total int
	for idx, value := range lefts {
		total += ints.IAbsDiff(value, rights[idx])
	}

	fmt.Println("Total distance : ", getTotalDistance(lefts, rights))
}

func part2() {
	lefts := []int{}
	rights := []int{}
	lines := inputreaders.ReadIntegersPerLines("./input.txt", "   ")
	for _, line := range lines {
		lefts = append(lefts, line[0])
		rights = append(rights, line[1])
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	fmt.Println("Total occurences : ", getTotalNumberOfOccurences(lefts, rights))
}

func main() {
	part1()
	part2()
}
