package day1

import "strings"

func extractDigit(line string) int {
	days := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	return len(days)
}

func PartB(data string) int {
	lines := strings.Split(string(data), "\n")

	total := 0

	for _, line := range lines {
		total += len(line) // dummay
	}

	return total
}
