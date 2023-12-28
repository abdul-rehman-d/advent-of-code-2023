package day4

import "strings"

func PartB(data string) int {
	lines := strings.Split(data, "\n")

	nums := make([]int, len(lines))

	for idx, line := range lines {
		nums[idx] += 1

		_, after, found := strings.Cut(line, ":")
		if !found {
			continue
		}
		winning, have, _ := strings.Cut(after, "|")
		winningNumbers, _ := parseIntArr(winning)
		haveNumbers, _ := parseIntArr(have)

		score := 0

		for num, _ := range winningNumbers {
			if haveNumbers[num] != 0 {
				score++
			}
		}

		for i := 1; i <= score; i++ {
			nums[idx+i] += nums[idx]
		}

	}

	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
