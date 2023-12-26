package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/abdul-rehman-d/advent-of-code-2023/day1"
	"github.com/abdul-rehman-d/advent-of-code-2023/day2"
)

func getChallengeInput(i int) string {
	arrByte, err := os.ReadFile(fmt.Sprintf("./day%d/input.txt", i))
	if err != nil {
		panic(err)
	}
	return string(arrByte)
}

func make(day int) {}

func solve(day int) {

	partA := []func(string) int{day1.PartA, day2.PartA}
	partB := []func(string) int{day1.PartB, day2.PartB}

	if day > len(partA) {
		log.Fatalf("Day cannot be more than %d\n", len(partA))
	}

	fmt.Printf("Part A Answer => %d\n", partA[day-1](getChallengeInput(day)))
	fmt.Printf("Part B Answer => %d\n", partB[day-1](getChallengeInput(day)))
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println(`Usage ./aoc [cmd] [day]
Available commands:-
- solve
- make`)
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error parsing day %s\n", os.Args[2])
		os.Exit(1)
	}

	if day <= 0 {
		fmt.Println("Day cannot be less than 1")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "solve":
		solve(day)
	case "make":
		make(day)
	default:
		fmt.Println("Invalid command.")
		os.Exit(1)
	}
}
