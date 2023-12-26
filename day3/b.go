package day3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func PartB(data string) int {
	lines := strings.Split(data, "\n")

	regex := regexp.MustCompile(`\*`)

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

	fmt.Println(len(coords))

	num := 0
	total := 0

	for _, coord := range coords {
		row := coord[0]
		col := coord[1]

		isLastCol := col >= len(lines[row])
		isFirstCol := col == 0

		isLastRow := row >= len(lines)
		isFirstRow := row == 0

		numOfAdjacents := 0
		product := 1

		if !isFirstRow {
			if !isFirstCol && allInts[row-1][col-1] != 0 {
				// top left
				product *= allInts[row-1][col-1]
				numOfAdjacents++
			}
			// top center
			if allInts[row-1][col] != 0 && allInts[row-1][col-1] != allInts[row-1][col] {
				product *= allInts[row-1][col]
				numOfAdjacents++
			}
			if !isLastCol && allInts[row-1][col+1] != 0 && allInts[row-1][col+1] != allInts[row-1][col] {
				// top right
				product *= allInts[row-1][col+1]
				numOfAdjacents++
			}
		}

		if !isFirstCol && allInts[row][col-1] != 0 {
			// current left
			product *= allInts[row][col-1]
			numOfAdjacents++
		}
		if !isLastCol && allInts[row][col+1] != 0 {
			// current right
			product *= allInts[row][col+1]
			numOfAdjacents++
		}

		if !isLastRow {
			if !isFirstCol && allInts[row+1][col-1] != 0 {
				// bottom left
				product *= allInts[row+1][col-1]
				numOfAdjacents++
			}
			// bottom center
			if allInts[row+1][col] != 0 && allInts[row+1][col-1] != allInts[row+1][col] {
				product *= allInts[row+1][col]
				numOfAdjacents++
			}
			if !isLastCol && allInts[row+1][col+1] != 0 && allInts[row+1][col+1] != allInts[row+1][col] {
				// bottom right
				product *= allInts[row+1][col+1]
				numOfAdjacents++
			}
		}

		if numOfAdjacents == 2 {
			num++
			fmt.Printf("Solution: %d\n", product)
			total += product
		}
	}

	fmt.Println(num)

	return total
}
