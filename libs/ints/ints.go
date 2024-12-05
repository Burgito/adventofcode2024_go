package ints

import "strconv"

func IAbsDiff(value1 int, value2 int) int {
	if value1 < value2 {
		return value2 - value1
	}

	return value1 - value2
}

func I64AbsDiff(value1 int64, value2 int64) int64 {
	if value1 < value2 {
		return value2 - value1
	}

	return value1 - value2
}

func ParseIntFromString(intString string) int {
	value, err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}

	return value
}
