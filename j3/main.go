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
	// multipliers := []int{}
	for _, line := range lines {
		//fmt.Printf("%q\n", re.FindAll([]byte(line), -1))
		for _, mul := range re.FindAllString(line, -1) {
			values := strings.ReplaceAll(strings.ReplaceAll(mul, "mul(", ""), ")", "")
			total += (ints.ParseIntFromString(strings.Split(values, ",")[0]) * ints.ParseIntFromString(strings.Split(values, ",")[1]))
		}
	}

	fmt.Println("Total of mul : ", total)
}

func part2() {
	fmt.Println("Part 2")
}

func main() {
	part1()
	part2()
}
