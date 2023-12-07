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
	// partOne(lines)
	partTwo(lines)
}

func partOne(text []string) {
	seeds := getSeeds(text[0])
	almanac := generateAlmanac(text)

	closestLocation := 0
	for i, s := range seeds {
		locationScore := score(s, almanac)
		if i == 0 {
			closestLocation = locationScore
		}
		if closestLocation > locationScore {
			closestLocation = locationScore
		}
	}

	fmt.Printf("closest location=%d\n", closestLocation)
}

func generateAlmanac(text []string) map[string][][]int {
	maps := make(map[string][][]int, 0)
	for _, t := range text {
		lines := strings.Split(t, "\n")
		hash := ""
		for i, line := range lines {
			if i == 0 {
				hash = line
				maps[line] = make([][]int, 0)
			} else {
				maps[hash] = append(maps[hash], getAlmanacRanges(line))
			}
		}
	}

	return maps
}

func getAlmanacRanges(s string) []int {
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

func score(seed int, maps map[string][][]int) int {
	soil := getDestination(seed, maps["seed-to-soil map:"])
	fertilizer := getDestination(soil, maps["soil-to-fertilizer map:"])
	water := getDestination(fertilizer, maps["fertilizer-to-water map:"])
	light := getDestination(water, maps["water-to-light map:"])
	temperature := getDestination(light, maps["light-to-temperature map:"])
	humidity := getDestination(temperature, maps["temperature-to-humidity map:"])
	location := getDestination(humidity, maps["humidity-to-location map:"])
	return location
}

func getDestination(source int, maps [][]int) int {
	for _, m := range maps {
		if source >= m[1] && source <= m[1]+m[2] {
			return m[0] + (source - m[1])
		}
	}

	return source
}

func partTwo(text []string) {
	seeds := getSeeds(text[0])
	almanac := generateAlmanac(text)
	fmt.Printf("seeds=%+v\n", seeds)
	fmt.Printf("almanac=%+v\n", almanac)

	closestLocation := 0
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j <= seeds[i]+seeds[i+1]-1; j++ {
			locationScore := score(j, almanac)
			if closestLocation == 0 {
				closestLocation = locationScore
				continue
			}

			if closestLocation > locationScore {
				closestLocation = locationScore
			}

		}
	}

	fmt.Printf("closest location=%d\n", closestLocation)
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
