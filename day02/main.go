package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Combiantion struct {
	Green int64
	Red   int64
	Blue  int64
}

func NewCombination() *Combiantion {
	c := &Combiantion{
		Blue:  0,
		Green: 0,
		Red:   0,
	}

	return c
}

var gameCombiantions = &Combiantion{
	Green: 13,
	Red:   12,
	Blue:  14,
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open(path + "/input.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var totalPartOne int64 = 0
	var totalPartTwo int64 = 0
	for fileScanner.Scan() {
		t := fileScanner.Text()
		v1 := scanLinePartOne(t)
		totalPartOne = totalPartOne + v1

		v2 := scanLinePartTwo(t)
		totalPartTwo = totalPartTwo + v2
	}

	fmt.Printf("total part one=%d\n", totalPartOne)
	fmt.Printf("total part two=%d\n", totalPartTwo)
}

func cleanValue(l string, old string) int64 {
	str := strings.TrimSpace(l)
	number := strings.Replace(str, old, "", 1)
	i, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func gameId(l string) int64 {
	return cleanValue(l, "Game ")
}

func getCombination(l string) *Combiantion {
	c := NewCombination()
	colors := strings.Split(l, ",")

	for _, color := range colors {
		if strings.HasSuffix(color, "blue") {
			v := cleanValue(color, " blue")
			c.Blue = v
		} else if strings.HasSuffix(color, "green") {
			v := cleanValue(color, " green")
			c.Green = v
		} else if strings.HasSuffix(color, "red") {
			v := cleanValue(color, " red")
			c.Red = v
		}
	}

	return c
}

func isCombinationPossible(c *Combiantion) bool {
	if c.Blue > gameCombiantions.Blue || c.Green > gameCombiantions.Green || c.Red > gameCombiantions.Red {
		return false
	}

	return true
}

func scanLinePartOne(l string) int64 {
	substrings := strings.Split(l, ":")
	bags := strings.Split(substrings[1], ";")
	for _, bag := range bags {
		parsedBag := getCombination(bag)
		if !isCombinationPossible(parsedBag) {
			return 0
		}
	}

	id := gameId(substrings[0])
	return id
}

func scanLinePartTwo(l string) int64 {
	substrings := strings.Split(l, ":")
	bags := strings.Split(substrings[1], ";")
	minCombination := NewCombination()
	for _, bag := range bags {
		parsedBag := getCombination(bag)
		if parsedBag.Blue > minCombination.Blue {
			minCombination.Blue = parsedBag.Blue
		}

		if parsedBag.Green > minCombination.Green {
			minCombination.Green = parsedBag.Green
		}

		if parsedBag.Red > minCombination.Red {
			minCombination.Red = parsedBag.Red
		}
	}

	value := minCombination.Blue * minCombination.Green * minCombination.Red
	return value
}
