package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInt64FromStr(text string) int64 {
	value, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		panic(err)
	}

	return value
}

func appendInt64(slice []int64, value int64) {
	slice = append(slice, value)
}

func i64AbsDiff(a int64, b int64) int64 {
	if a < b {
		return b - a
	}

	return a - b
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	lefts := []int64{}
	rights := []int64{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), "   ")
		left := parseInt64FromStr(pair[0])
		right := parseInt64FromStr(pair[1])

		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	var total int64
	for idx, value := range lefts {
		total += int64(i64AbsDiff(value, rights[idx]))
	}

	fmt.Println(total)
}
