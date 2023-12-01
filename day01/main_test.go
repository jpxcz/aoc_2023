package main

import "testing"

func TestPartOneTotal(t *testing.T) {
	lines := make([]string, 4)
	lines[0] = "1abc2"
	lines[1] = "pqr3stu8vwx"
	lines[2] = "a1b2c3d4e5f"
	lines[3] = "treb7uchet"
	file := struct {
		lines []string
		total int
	}{
		lines: lines,
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
	lines := make([]string, 7)
	lines[0] = "two1nine"
	lines[1] = "eightwothree"
	lines[2] = "abcone2threexyz"
	lines[3] = "xtwone3four"
	lines[4] = "4nineeightseven2"
	lines[5] = "zoneight234"
	lines[6] = "7pqrstsixteen"
	file := struct {
		lines []string
		total int
	}{
		lines: lines,
		total: 281,
	}

	for _, l := range file.lines {
		calculatePartTwo(l)
	}

	if totalPartTwo != file.total {
		t.Fatalf("expected=%d, got=%d", file.total, totalPartTwo)
	}
}