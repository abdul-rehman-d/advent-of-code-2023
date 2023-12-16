package day2

import (
	"regexp"
	"strconv"
	"strings"
)

type Set = map[string]int
type Game = map[int]([]Set)

func getGame(line string) Game {
	intRegex := regexp.MustCompile(`\d+`)
	gameId := intRegex.FindString(line)
	parsedGameId, _ := strconv.Atoi(gameId)

	out := make(Game)
	out[parsedGameId] = []Set{}

	sets := strings.Split(strings.Split(line, ": ")[1], "; ")

	for i := 0; i < len(sets); i++ {
		set := sets[i]
		setObj := make(Set)
		balls := strings.Split(set, ", ")
		for j := 0; j < len(balls); j++ {
			ball := strings.Split(balls[j], " ")
			val, _ := strconv.Atoi(ball[0])

			setObj[ball[1]] = val
		}
		out[parsedGameId] = append(out[parsedGameId], setObj)
	}
	return out
}

func PartA(data string) int {
	lines := strings.Split(data, "\n")

	sum_of_ids := 0

	possible := make(map[string]int)
	possible["red"] = 12
	possible["green"] = 13
	possible["blue"] = 14

	for _, line := range lines {
		game := getGame(line)
		for gameId, sets := range game {
			is_possible := true
			for _, set := range sets {
				for color, amount := range set {
					if possible[color] < amount {
						is_possible = false
						break
					}
				}
				if !is_possible {
					break
				}
			}
			if is_possible {
				sum_of_ids += gameId
			}
		}
	}

	return sum_of_ids
}
