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

	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		if isDigit(char) {
			first = char
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		char := rune(s[i])
		if isDigit(char) {
			second = char
			break
		}
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
