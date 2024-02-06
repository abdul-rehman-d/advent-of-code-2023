package day5

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parseSeedsB(str string) []SeedRange {
	after, _ := strings.CutPrefix(str, "seeds: ")
	seedsStr := strings.Split(after, " ")

	parsed := make([]SeedRange, len(seedsStr)/2)

	idx := 0

	for i := 0; i < len(seedsStr); i += 2 {
		start, _ := strconv.Atoi(seedsStr[i])
		count, _ := strconv.Atoi(seedsStr[i+1])

		parsed[idx] = SeedRange{
			start,
			count,
		}
		idx++
	}

	return parsed
}

type SeedRange struct {
	start int
	count int
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func getMappedValueB(curr SeedRange, rows Rows) SeedRange {
	for _, currentRow := range rows {
		if currentRow.source <= curr.start && currentRow.source+currentRow.rnge > curr.start {
			return SeedRange{
				start: currentRow.destination + (curr.start - currentRow.source),
				count: min(currentRow.rnge-curr.start-currentRow.source, curr.count),
			}
		}
	}
	return curr
}

func PartB(data string) int {
	splitted := strings.Split(data, "\n\n")

	// first row: seeds
	currentMappedValues := parseSeedsB(splitted[0])
	fmt.Println(currentMappedValues)
	// rest: maps
	allMaps := parseMaps(splitted[1:])

	for _, currentMap := range allMaps {
		for i := 0; i < len(currentMappedValues); i++ {
			currentMappedValues[i] = getMappedValueB(currentMappedValues[i], currentMap)
		}
		fmt.Println(currentMappedValues)
	}

	lowest := math.MaxInt64
	for i := 0; i < len(currentMappedValues); i++ {
		if lowest > currentMappedValues[i].start {
			lowest = currentMappedValues[i].start
		}
	}

	return lowest
}
