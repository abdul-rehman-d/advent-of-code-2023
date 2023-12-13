package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abdul-rehman-d/advent-of-code-2023/day1"
)

func getChallengeInput(i int) string {
	arrByte, err := os.ReadFile(fmt.Sprintf("./day%d/input.txt", i))
	if err != nil {
		panic(err)
	}
	return string(arrByte)
}

func main() {

	partA := []func(string) int{day1.PartA}
	partB := []func(string) int{day1.PartB}

	fmt.Print("Enter the day you want to run: ")

	var day int
	_, err := fmt.Scan(&day)

	if err != nil {
		panic(err)
	}

	if day <= 0 {
		log.Fatalln("Day cannot be less than 1")
	}

	if day > len(partA) {
		log.Fatalf("Day cannot be more than %d\n", len(partA))
	}

	fmt.Printf("Part A Answer => %d\n", partA[day-1](getChallengeInput(day)))
	fmt.Printf("Part B Answer => %d\n", partB[day-1](getChallengeInput(day)))
}
