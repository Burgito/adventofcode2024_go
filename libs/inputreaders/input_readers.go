package inputreaders

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func AppendValuesFromLine(intsLine []int, line []string) []int {
	for _, value := range line {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		intsLine = append(intsLine, intValue)
	}

	return intsLine
}

func ReadLinesAsStrings(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadFullFileAsString(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	fileContent := ""
	for scanner.Scan() {
		fileContent += scanner.Text()
	}

	return fileContent
}

func ReadIntegersPerLines(filePath string, separator string) [][]int {
	lines := ReadLinesAsStrings(filePath)
	intsLines := [][]int{}
	intsLine := []int{}
	for _, line := range lines {
		intsLines = append(intsLines, AppendValuesFromLine(intsLine, strings.Split(line, separator)))
	}

	return intsLines
}
