package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TextScan []string

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open(path + "/input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	text := make(TextScan, 0)

	for fileScanner.Scan() {
		t := fileScanner.Text()
		text = append(text, t)
	}

	partOne(text)
	partTwo(text)
}

type NumberCoordinates struct {
	value  int
	x      int
	y      int
	yEnd   int
	len    int
	parsed bool
}

func NewNumberCoordinates(x, y, yEnd int, numberString string) *NumberCoordinates {
	n, err := strconv.Atoi(numberString)
	if err != nil {
		panic(err)
	}

	return &NumberCoordinates{
		value:  n,
		x:      x,
		y:      y,
		yEnd:   yEnd,
		len:    len(numberString),
		parsed: false,
	}
}

func getNumberCoordinates(text TextScan) []*NumberCoordinates {
	numbers := make([]*NumberCoordinates, 0)

	for lIndex, line := range text {
		numberStart := -1
		numberEnd := -1
		for rIndex, r := range line {
			if r >= '0' && r <= '9' {
				if numberStart == -1 {
					numberStart = rIndex
					numberEnd = rIndex
				} else {
					numberEnd = rIndex
				}
			} else {
				if numberStart != -1 {
					n := NewNumberCoordinates(lIndex, numberStart, numberEnd, line[numberStart:numberEnd+1])
					numbers = append(numbers, n)
				}
				numberStart = -1
				numberEnd = -1
			}
		}

		if numberStart != -1 {
			n := NewNumberCoordinates(lIndex, numberStart, numberEnd, line[numberStart:numberEnd+1])
			numbers = append(numbers, n)
		}
	}

	return numbers
}

func isSymbol(r rune) bool {
	if r >= '0' && r <= '9' || r == '.' {
		return false
	}

	return true
}

func isAsterik(r rune) bool {
	return r == '*'
}

func findAdjacentNumbers(x, y int, numbers []*NumberCoordinates) int {
	adjacents := 0
	for _, n := range numbers {
		if n.parsed {
			continue
		}

		if n.x >= x-1 && n.x <= x+1 && ((n.y >= y-1 && n.y <= y+1) || (n.yEnd >= y-1 && n.yEnd <= y+1)) {
			adjacents = adjacents + n.value
			n.parsed = true
		}
	}

	return adjacents
}

func findAdjacentSymbolsSumatory(text TextScan, numbers []*NumberCoordinates) int {
	total := 0
	for lIndex, line := range text {
		for rIndex, r := range line {
			if isSymbol(r) {
				total = total + findAdjacentNumbers(lIndex, rIndex, numbers)
			}
		}
	}

	return total
}

func partOne(text TextScan) {
	numbers := getNumberCoordinates(text)
	total := findAdjacentSymbolsSumatory(text, numbers)
	fmt.Printf("total part one=%d\n", total)
}

func findAdjacentGearNumbers(x, y int, numbers []*NumberCoordinates) int {
	firstNumber := 0
	secondNumber := 0
	for _, n := range numbers {
		if n.parsed {
			continue
		}

		if n.x >= x-1 && n.x <= x+1 && ((n.y >= y-1 && n.y <= y+1) || (n.yEnd >= y-1 && n.yEnd <= y+1)) {
			if firstNumber == 0 {
				firstNumber = n.value
			} else {
				secondNumber = n.value
			}
		}
	}

	return firstNumber * secondNumber

}

func findGearsSumatory(text TextScan, numbers []*NumberCoordinates) int {
	total := 0

	for lIndex, line := range text {
		for rIndex, r := range line {
			if isAsterik(r) {
				total = total + findAdjacentGearNumbers(lIndex, rIndex, numbers)
			}
		}
	}

	return total
}

func partTwo(text TextScan) {
	numbers := getNumberCoordinates(text)
	total := findGearsSumatory(text, numbers)
	fmt.Printf("total part two=%d\n", total)
}
