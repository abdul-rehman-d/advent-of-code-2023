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

func getFirstDigit(s string) int {
	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		if isDigit(char) {
			val, _ := strconv.Atoi(string(char))
			return val
		}
	}

	return 0
}

func getLastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		char := rune(s[i])
		if isDigit(char) {
			val, _ := strconv.Atoi(string(char))
			return val
		}
	}

	return 0
}

func getTwoDigits(s string) int {
	first := getFirstDigit(s)
	second := getLastDigit(s)

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
