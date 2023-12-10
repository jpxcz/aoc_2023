package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	HIGH_CARD       int = 1
	ONE_PAIR            = 2
	TWO_PAIRS           = 3
	THREE_OF_A_KIND     = 4
	FULL_HOUSE          = 5
	FOUR_OF_A_KIND      = 6
	FIVE_OF_A_KIND      = 7
)

var cardStrenght = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 0,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'1': 1,
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

	text := make([]string, 0)

	for fileScanner.Scan() {
		t := fileScanner.Text()
		text = append(text, t)
	}

	partTwo(text)
}

// type ScartchCardsCopies map[int]int

type CamelHand struct {
	values   string
	handType int
	bid      int
}

func NewCamelHand(values string, handType int, bid int) *CamelHand {
	return &CamelHand{
		values:   values,
		handType: handType,
		bid:      bid,
	}
}

func partTwo(text []string) {
	cards := make([]*CamelHand, 0)
	for _, line := range text {
		values := strings.Split(line, " ")
		bid := getNumberValue(values[1])
		hand := values[0]
		handValue := getHandValue(hand)
		camelHand := NewCamelHand(hand, handValue, bid)
		cards = append(cards, camelHand)
	}

	total := 0
	var currentLowest *CamelHand = nil
	indexLowest := -1
	cardLenght := len(cards)
	for rank := 1; rank <= cardLenght; rank++ {
		for idx, c := range cards {
			currentLowest = getLowestCard(currentLowest, c)
			if currentLowest == c {
				indexLowest = idx
			}
		}
		total = total + (currentLowest.bid * rank)
		cards = remove(cards, indexLowest)
		currentLowest = nil
		indexLowest = -1
	}

	fmt.Printf("total part two=%d\n", total)
}

func getNumberValue(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

func getHandValue(s string) int {
	m := map[rune]int{}

	for _, r := range s {
		m[r] += 1
	}

	jokers := 0
	if v, ok := m['J']; ok {
		jokers = v
		delete(m, 'J')
	}

	highestCard := 'J'
	highestValue := 0
	for k, v := range m {
		if v > highestValue {
			highestCard = k
			highestValue = v
		}
	}

	m[highestCard] = m[highestCard] + jokers

	if len(m) == 1 {
		return FIVE_OF_A_KIND
	}

	if len(m) == 2 {
		for _, v := range m {
			if v == 4 {
				return FOUR_OF_A_KIND
			}
		}
		return FULL_HOUSE
	}

	if len(m) == 3 {
		for _, v := range m {
			if v == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIRS
	}

	if len(m) == 4 {
		return ONE_PAIR
	}

	return HIGH_CARD
}

func isHandOneBetter(handOne, handTwo string) bool {
	for i, _ := range handOne {
		if cardStrenght[handOne[i]] == cardStrenght[handTwo[i]] {
			continue
		}

		if cardStrenght[handOne[i]] > cardStrenght[handTwo[i]] {
			return true
		}

		return false
	}

	return false
}

func remove(slice []*CamelHand, s int) []*CamelHand {
	return append(slice[:s], slice[s+1:]...)
}

func getLowestCard(currLowest, newCard *CamelHand) *CamelHand {
	if currLowest == nil {
		return newCard
	}

	if currLowest.handType > newCard.handType {
		return newCard
	}

	if currLowest.handType < newCard.handType {
		return currLowest
	}

	handOneBetter := isHandOneBetter(currLowest.values, newCard.values)
	if handOneBetter {
		return newCard
	} else {
		return currLowest
	}
}
