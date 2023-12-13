package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func isDigit(char rune) bool {
	ascii := int(char)
	return ascii < 58 && ascii > 47
}

func getTwoDigits(s string) int {
	var first rune
	var second rune

	for _, char := range s {
		if isDigit(char) {
			if first != 0 {
				second = char
			} else {
				first = char
			}
		}
	}

	if second == 0 {
		second = first
	}

	val, _ := strconv.Atoi(fmt.Sprintf("%c%c", first, second))
	return val
}

func PartA(data string) int {
	lines := strings.Split(string(data), "\n")

	total := 0

	for _, line := range lines {
		total += getTwoDigits(line)
	}

	return total
}
