package main

import (
	"adventofcode/inputreaders"
	"adventofcode/ints"
	"fmt"
	"regexp"
	"strings"
)

func part1() {
	lines := inputreaders.ReadLinesAsStrings("./input.txt")
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	total := 0
	for _, line := range lines {
		for _, mul := range re.FindAllString(line, -1) {
			values := strings.ReplaceAll(strings.ReplaceAll(mul, "mul(", ""), ")", "")
			total += (ints.ParseIntFromString(strings.Split(values, ",")[0]) * ints.ParseIntFromString(strings.Split(values, ",")[1]))
		}
	}

	fmt.Println("Total of mul : ", total)
}

func part2() {
	memory := inputreaders.ReadFullFileAsString("./input.txt")
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	total := 0

	shouldMultiply := true
	for _, match := range re.FindAllString(memory, -1) {

		if !shouldMultiply {
			if match == "do()" {
				shouldMultiply = true
			}
			continue
		}

		if match == "don't()" {
			shouldMultiply = false
			continue
		}

		if match == "do()" {
			continue
		}

		values := strings.ReplaceAll(strings.ReplaceAll(match, "mul(", ""), ")", "")
		total += (ints.ParseIntFromString(strings.Split(values, ",")[0]) * ints.ParseIntFromString(strings.Split(values, ",")[1]))
	}

	fmt.Println("New total : ", total)
}

func main() {
	part1()
	part2()
}
