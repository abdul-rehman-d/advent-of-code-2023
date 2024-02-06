package day5

import (
	"math"
	"strconv"
	"strings"
)

type IntArr []int

func parseSeeds(str string) []int {
	after, _ := strings.CutPrefix(str, "seeds: ")
	seedsStr := strings.Split(after, " ")

	parsed := make([]int, len(seedsStr))

	for i := 0; i < len(seedsStr); i++ {
		parsed[i], _ = strconv.Atoi(seedsStr[i])
	}

	return parsed
}

type Row struct {
	source      int
	destination int
	rnge        int
}

type Rows []Row

func parseMaps(allMaps []string) []Rows {
	out := make([]Rows, len(allMaps))

	for currentMapIdx, currentMapRaw := range allMaps {
		currentMapLines := strings.Split(currentMapRaw, "\n")

		out[currentMapIdx] = make(Rows, len(currentMapLines)-1)

		for currentRowIdx := 1; currentRowIdx < len(currentMapLines); currentRowIdx++ {
			currentRowRaw := currentMapLines[currentRowIdx]
			currentRowNumbers := strings.Split(currentRowRaw, " ")

			destination, _ := strconv.Atoi(currentRowNumbers[0])
			source, _ := strconv.Atoi(currentRowNumbers[1])
			rnge, _ := strconv.Atoi(currentRowNumbers[2])

			out[currentMapIdx][currentRowIdx-1] = Row{
				source:      source,
				destination: destination,
				rnge:        rnge,
			}
		}
	}

	return out
}

func getMappedValue(curr int, rows Rows) int {
	for _, currentRow := range rows {
		if currentRow.source <= curr && currentRow.source+currentRow.rnge > curr {
			return currentRow.source + (curr - currentRow.source) + (currentRow.destination - currentRow.source)
		}
	}
	return curr
}

func PartA(data string) int {
	splitted := strings.Split(data, "\n\n")

	// first row: seeds
	currentMappedValues := parseSeeds(splitted[0])
	// rest: maps
	allMaps := parseMaps(splitted[1:])

	for _, currentMap := range allMaps {
		for i := 0; i < len(currentMappedValues); i++ {
			currentMappedValues[i] = getMappedValue(currentMappedValues[i], currentMap)
		}
	}

	lowest := math.MaxInt64
	for i := 0; i < len(currentMappedValues); i++ {
		if lowest > currentMappedValues[i] {
			lowest = currentMappedValues[i]
		}
	}

	return lowest
}
