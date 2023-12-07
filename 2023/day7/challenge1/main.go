package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardToNumber = map[byte]int{
	'2':  2,
	'3':  3,
	'4':  4,
	'5':  5,
	'6':  6,
	'7':  7,
	'8':  8,
	'9':  9,
	'T':  10,
	'J':  11,
	'Q':  12,
	'K':  13,
	'A':  14,
}

type hand struct {
	cards string
	bid	 int
}

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := []hand{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		bid, err := strconv.Atoi(line[6:])

		if err != nil {
			log.Fatal(err)
		}

		hands = append(hands, hand{
			cards: line[:5],
			bid: bid,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i].cards, hands[j].cards)
	})

	totalWinnings := 0
	for i, hand := range hands {
		fmt.Println(hand.cards)
		totalWinnings += (i+1) * hand.bid
	}

	fmt.Println(totalWinnings)
}

func compareHands(hand1 string, hand2 string) bool {
	hand1Score := handToScore(hand1)
	hand2Score := handToScore(hand2)

	if hand1Score != hand2Score {
		return hand1Score < hand2Score
	} else {
		for i := 0; i < len(hand1); i++ {
			if hand1[i] != hand2[i] {
				return cardToNumber[hand1[i]] < cardToNumber[hand2[i]]
			}
		}

		log.Fatal("Hands are equal")
		return true
	}
}

func handToScore(hand string) int {
	numberOfEachCard := map[string]int{}

	for _, card := range hand {
		numberOfEachCard[string(card)]++
	}

	numberToCards := map[int]int{}
	for _, number := range numberOfEachCard {
		numberToCards[number]++
	}

	if numberToCards[5] == 1 {
		return 1_000_000
	} else if numberToCards[4] == 1 {
		return 100_000
	} else if numberToCards[3] == 1 && numberToCards[2] == 1 {
		return 10_000
	} else if numberToCards[3] == 1 {
		return 1_000
	} else if numberToCards[2] == 2 {
		return 100
	} else if numberToCards[2] == 1 {
		return 10
	} else {
		return 1
	}
}


