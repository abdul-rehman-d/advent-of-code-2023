package day1

import (
	"testing"
)

func TestPartA(t *testing.T) {
	data := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	expected := 142
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	data := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	expected := 281
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}
