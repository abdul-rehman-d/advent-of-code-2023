package day1

import (
	"strconv"
	"strings"
)

func isDigit(char rune) bool {
	ascii := int(char)
	return ascii < 58 && ascii > 47
}

func getFirstDigit(s string) (int, int) {
	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		if isDigit(char) {
			val, _ := strconv.Atoi(string(char))
			return i, val
		}
	}

	return -1, 0
}

func getLastDigit(s string) (int, int) {
	for i := len(s) - 1; i >= 0; i-- {
		char := rune(s[i])
		if isDigit(char) {
			val, _ := strconv.Atoi(string(char))
			return i, val
		}
	}

	return -1, 0
}

func PartA(data string) int {
	lines := strings.Split(string(data), "\n")

	total := 0

	for _, line := range lines {
		_, first := getFirstDigit(line)
		_, second := getLastDigit(line)
		total += ((first * 10) + second)
	}

	return total
}
