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

	lines := strings.Split(string(content), "\n")
	partOne(lines)

}

func partOne(text []string) {
	times := toIntegers(strings.Fields(text[0])[1:])
	distances := toIntegers(strings.Fields(text[1])[1:])

	beats := 0
	for i := 0; i < len(times); i++ {
		wons := countPassingOptions(times[i], distances[i])
		if i == 0 {
			beats = wons
			continue
		}

		beats = wons * beats
	}

	fmt.Printf("total=%d\n", beats)
}

func toIntegers(values []string) []int {
	numbers := make([]int, 0)
	for _, v := range values {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, n)
	}

	return numbers
}

func countPassingOptions(time int, distance int) int {
	count := 0
	for i := 1; i <= time; i++ {
		if (i * (time-i)) > distance {
			count += 1
		}
	}

	return count
}