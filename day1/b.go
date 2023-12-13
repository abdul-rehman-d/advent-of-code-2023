package day1

import (
	"strings"
)

func getFirstNamedDigit(days map[string]int, line string) (int, int) {
	for i := 0; i < len(line)-3; i++ {
		for j := 3; j <= 5; j++ {
			if i+j > len(line) {
				continue
			}
			day := line[i : i+j]
			if _, ok := days[day]; ok {
				return i, days[day]
			}
		}
	}
	return -1, 0
}

func getLastNamedDigit(days map[string]int, line string) (int, int) {
	for i := len(line) - 3; i >= 0; i-- {
		for j := 3; j <= 5; j++ {
			if i+j > len(line) {
				continue
			}
			day := line[i : i+j]
			if _, ok := days[day]; ok {
				return i, days[day]
			}
		}
	}
	return -1, 0
}

func getMin(idx1, f1, idx2, f2 int) (int, int) {
	if idx1 == -1 {
		return idx2, f2
	}
	if idx2 == -1 {
		return idx1, f1
	}
	if idx1 < idx2 {
		return idx1, f1
	}
	return idx2, f2
}

func getMax(idx1, f1, idx2, f2 int) (int, int) {
	if idx1 == -1 {
		return idx2, f2
	}
	if idx2 == -1 {
		return idx1, f1
	}
	if idx1 > idx2 {
		return idx1, f1
	}
	return idx2, f2
}

func PartB(data string) int {
	days := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	lines := strings.Split(string(data), "\n")

	total := 0

	for _, line := range lines {
		firstIdx, first := getFirstDigit(line)
		lastIdx, last := getLastDigit(line)

		firstNamedIdx, firstNamed := getFirstNamedDigit(days, line)
		lastNamedIdx, lastNamed := getLastNamedDigit(days, line)

		_, trueFirst := getMin(firstIdx, first, firstNamedIdx, firstNamed)
		_, trueLast := getMax(lastIdx, last, lastNamedIdx, lastNamed)

		total += ((trueFirst * 10) + trueLast)
	}

	return total
}
