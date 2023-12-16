package day2

import "strings"

func PartB(data string) int {
	lines := strings.Split(data, "\n")

	sum := 0

	for _, line := range lines {
		game := getGame(line)
		for _, sets := range game {
			min := make(map[string]int, 2)
			min["red"] = 0
			min["green"] = 0
			min["blue"] = 0
			for _, set := range sets {
				for color, amount := range set {
					if amount > min[color] {
						min[color] = amount
					}
				}
			}
			sum += (min["red"] * min["green"] * min["blue"])
		}
	}
	return sum
}
