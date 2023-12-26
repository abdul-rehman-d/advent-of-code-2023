package day3

import (
	"regexp"
	"strconv"
	"strings"
)

func isDigit(char rune) bool {
	ascii := int(char)
	return ascii < 58 && ascii > 47
}

func PartA(data string) int {
	lines := strings.Split(data, "\n")

	regex := regexp.MustCompile(`[^\w.]`)

	allInts := make([][]int, len(lines))
	for i := range allInts {
		allInts[i] = make([]int, len(lines[0]))
	}

	coords := [][]int{}

	for y, line := range lines {
		digitYet := ""

		for x, char := range line {
			allInts[y][x] = 0

			if isDigit(char) {
				digitYet += string(char)
			} else {
				// int has ended
				if len(digitYet) > 0 {
					val, err := strconv.Atoi(digitYet)
					if err != nil {
						panic(err)
					}
					for i := 0; i < len(digitYet); i++ {
						allInts[y][x-i-1] = val
					}
					digitYet = ""
				}
			}
		}

		if len(digitYet) > 0 {
			val, err := strconv.Atoi(digitYet)
			if err != nil {
				panic(err)
			}
			for i := 0; i < len(digitYet); i++ {
				allInts[y][len(line)-i-1] = val
			}
			digitYet = ""
		}

		for _, coord := range regex.FindAllStringIndex(line, -1) {
			x := coord[0]
			coords = append(coords, []int{y, x})
		}
	}

	total := 0

	for _, coord := range coords {
		row := coord[0]
		col := coord[1]

		isLastCol := col >= len(lines[row])
		isFirstCol := col == 0

		isLastRow := row >= len(lines)
		isFirstRow := row == 0

		if !isFirstRow {
			if !isFirstCol {
				// top left
				total += allInts[row-1][col-1]
			}
			// top center
			if allInts[row-1][col-1] != allInts[row-1][col] {
				total += allInts[row-1][col]
			}
			if !isLastCol && allInts[row-1][col+1] != allInts[row-1][col] {
				// top right
				total += allInts[row-1][col+1]
			}
		}

		if !isFirstCol {
			// current left
			total += allInts[row][col-1]
		}
		if !isLastCol {
			// current right
			total += allInts[row][col+1]
		}

		if !isLastRow {
			if !isFirstCol {
				// bottom left
				total += allInts[row+1][col-1]
			}
			// bottom center
			if allInts[row+1][col-1] != allInts[row+1][col] {
				total += allInts[row+1][col]
			}
			if !isLastCol && allInts[row+1][col+1] != allInts[row+1][col] {
				// bottom right
				total += allInts[row+1][col+1]
			}
		}

	}

	return total
}
