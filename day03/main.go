package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

}

func partOne(text TextScan) {
	total := 0
	for lIndex, line := range text {
		initRange := 0
		lastRange := 0
		digit := false
		for tIndex, t := range line {
			if unicode.IsDigit(t) {
				if !digit {
					initRange = tIndex
				}
				digit = true
				lastRange = tIndex + 1
			} else {
				if digit {
					total = total + validateMatrix(initRange, lastRange, text, lIndex)
				}

				digit = false
				initRange = tIndex
				lastRange = tIndex
			}

			if digit && tIndex == len(line) -1 {
					total = total + validateMatrix(initRange, lastRange, text, lIndex)
			}
		}
	}

	fmt.Printf("total=%d\n", total)
}

func hasSymbol(chars string) bool {
	return strings.ContainsAny(chars, "@#$%^&*()-_=+[]|/")
}

func validateMatrix(initialRange int, finalRange int, text TextScan, currLine int) int {
	number := text[currLine][initialRange:finalRange]
	if initialRange > 0 {
		initialRange = initialRange - 1
	}

	if finalRange < len(text[0])-2 {
		finalRange = finalRange + 1

	}

	for i := currLine - 1; i <= currLine+1; i++ {
		if i < 0 || i >= len(text) {
			continue
		}
		fmt.Println(i)
		if hasSymbol(text[i][initialRange:finalRange]) {
			fmt.Printf("valid number %v\n", number)
			intVar, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			return intVar
		}
	}

	return 0
}
