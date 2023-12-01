package main

import "testing"

func TestPartOneTotal(t *testing.T) {
	file := struct {
		lines []string
		total int
	}{
		lines: []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
		total: 142,
	}

	for _, l := range file.lines {
		calculatePartOne(l)
	}

	if totalPartOne != file.total {
		t.Fatalf("expected=%d, got=%d", file.total, totalPartOne)
	}
}

func TestPartTwoTotal(t *testing.T) {
	file := struct {
		lines []string
		total int
	}{
		lines: []string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
		},
		total: 281,
	}

	for _, l := range file.lines {
		calculatePartTwo(l)
	}

	if totalPartTwo != file.total {
		t.Fatalf("expected=%d, got=%d", file.total, totalPartTwo)
	}
}
