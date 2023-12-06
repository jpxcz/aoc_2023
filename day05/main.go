package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(path + "/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n\n")
	partOne(lines)
}

func partOne(text []string) {
	seeds := getSeeds(text[0])
	closesLocation := int(999999999999999999)
	bestSeed := int(999999999999999999)
	for _, s := range seeds {
		p := s
		for i := 1; i < len(text); i++ {
			newProperty := getProperty(p, text[i])
			p = newProperty
		}
		if closesLocation > p {
			closesLocation = p
			bestSeed = s
		}
	}

	fmt.Printf("closest seed=%d, location=%d", closesLocation, bestSeed)
}

func getSeeds(l string) []int {
	split := strings.Split(l, "seeds:")
	sn := strings.Fields(split[1])

	numbers := make([]int, 0)
	for _, n := range sn {
		number, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
	}
	return numbers
}

func getProperty(initial int, text string) int {
	lines := strings.Split(text, "\n")
	for i := 1; i < len(lines); i++ {
		values := getValueProperties(lines[i])
		if initial >= values[1] && initial <= values[1]+values[2]-1 {
			dv := initial - values[1] + values[0]
			return dv
		}
	}

	return initial
}

func getValueProperties(s string) []int {
	sv := strings.Fields(s)

	numbers := make([]int, 0)
	for _, n := range sv {
		number, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
	}
	return numbers
}
