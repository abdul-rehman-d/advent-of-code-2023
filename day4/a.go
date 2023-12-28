package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func parseIntArr(s string) (map[int]int, error) {
	out := make(map[int]int)

	for _, val := range strings.Fields(s) {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		out[parsed] = parsed
	}

	return out, nil
}

func PartA(data string) int {
	lines := strings.Split(data, "\n")

	total := 0

	for _, line := range lines {
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
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		total += score
	}
	return total
}
