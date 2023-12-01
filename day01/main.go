package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)


var totalPartOne = 0
var totalPartTwo = 0

// part one regex solution -- first attempt created
// func partOneLineNumber(s string) int {
// 	re := regexp.MustCompile("[0-9]+")
// 	values := strings.Join(re.FindAllString(s, -1), "")
// 	if len(values) == 0 {
// 		return 0
// 	}

// 	str := string(values[0]) + string(values[len(values)-1])
// 	n, _ := strconv.Atoi(str)
// 	return n
// }

func firstDigit(runes []rune) int {
	for i := 0; i < len(runes); i++ {
		if unicode.IsNumber(runes[i]) {
			buf := make([]byte, 1)
			_ = utf8.EncodeRune(buf, runes[i])
			value, _ := strconv.Atoi(string(buf))
			return value
		}
	}

	return 0
}

func lastDigit(runes []rune) int {
	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsNumber(runes[i]) {
			buf := make([]byte, 1)
			_ = utf8.EncodeRune(buf, runes[i])
			value, _ := strconv.Atoi(string(buf))
			return value
		}
	}
	return 0
}

func replaceStrings(initialString string) string {
	s := strings.ReplaceAll(initialString, "zero", "0o")
	s = strings.ReplaceAll(s, "one", "o1e")
	s = strings.ReplaceAll(s, "two", "t2o")
	s = strings.ReplaceAll(s, "three", "e3e")
	s = strings.ReplaceAll(s, "four", "f4r")
	s = strings.ReplaceAll(s, "five", "f5e")
	s = strings.ReplaceAll(s, "six", "s6x")
	s = strings.ReplaceAll(s, "seven", "s7n")
	s = strings.ReplaceAll(s, "eight", "e8t")
	s = strings.ReplaceAll(s, "nine", "n9e")
	return s
}


func calculateNumbers(runes []rune) int {
	initialNumber := firstDigit(runes)
	secondNumber := lastDigit(runes)
	return initialNumber*10 + secondNumber
}

func calculatePartOne(s string) {
	runes := []rune(s)
	number := calculateNumbers(runes)
	totalPartOne = totalPartOne + number
}

func calculatePartTwo(s string) {
	rs := replaceStrings(s)
	runes := []rune(rs)
	number := calculateNumbers(runes)
	totalPartTwo = totalPartTwo + number 
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

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		t := fileScanner.Text()
		calculatePartOne(t)
		calculatePartTwo(t)
	}

	readFile.Close()
	fmt.Printf("total part one value is: %d\n", totalPartOne)
	fmt.Printf("total part two value is: %d\n", totalPartTwo)
}
