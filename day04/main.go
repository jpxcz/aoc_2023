package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

type CardGame struct {
	id             int
	hand           []int
	winningNumbers map[int]bool
}

func NewCardGame(id int, hand []int, winningNumbers map[int]bool) *CardGame {
	return &CardGame{
		id:             id,
		hand:           hand,
		winningNumbers: winningNumbers,
	}
}

func (c *CardGame) points() int {
	points := 0
	for _, n := range c.hand {
		if _, ok := c.winningNumbers[n]; ok {
			if points == 0 {
				points = 1
			} else {
				points = points * 2
			}
		}
	}

	return points
}

func partOne(text TextScan) {
	total := 0
	for _, line := range text {
		cardHands := strings.Split(line, ":")
		id := getCardId(cardHands[0])
		hands := strings.Split(cardHands[1], "|")
		hand := getCardHand(hands[0])
		winning := getWinningNumbers(hands[1])
		card := NewCardGame(id, hand, winning)
		total += card.points()

	}

	fmt.Printf("total part one=%d\n", total)
}

func sanitizeValue(s string, removeText string) int {
	sanitizedString := strings.TrimSpace(strings.ReplaceAll(s, removeText, ""))
	number, err := strconv.Atoi(sanitizedString)
	if err != nil {
		panic(err)
	}

	return number
}

func getCardId(s string) int {
	return sanitizeValue(s, "Card")
}

func getCardHand(s string) []int {
	hand := strings.Fields(s)
	handNumbers := make([]int, 0)
	for _, h := range hand {
		hn := sanitizeValue(h, "")
		handNumbers = append(handNumbers, hn)
	}

	return handNumbers
}

func getWinningNumbers(s string) map[int]bool {
	sw := strings.Fields(s)
	winning := make(map[int]bool)
	for _, sn := range sw {
		n := sanitizeValue(sn, "")
		winning[n] = true
	}

	return winning
}
