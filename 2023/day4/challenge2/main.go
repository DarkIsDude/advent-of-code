package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	numberOfCards := map[int]int{}

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ": ")
		cardAsString := strings.Replace(split[0], "Card ", "", -1)
		card, err := strconv.Atoi(strings.TrimSpace(cardAsString))

		if err != nil {
			log.Fatal(err)
		}

		numberOfCards[card] = numberOfCards[card] + 1
		lotery := strings.Split(strings.TrimSpace((split[1])), " | ")

		winning := strings.Split(lotery[0], " ")
		numbers := strings.Split(lotery[1], " ")

		numberOfWinning := 0

		for _, number := range numbers {
			for _, win := range winning {
				if number == win && number != "" && win != "" {
					numberOfWinning++
				}
			}
		}

		for i := 1; i <= numberOfWinning; i++ {
			numberOfCards[card + i] += numberOfCards[card]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for key, value := range numberOfCards {
		sum += value
		fmt.Printf("Card %d: %d\n", key, value)
	}

	fmt.Printf("Total: %d\n", sum)
}
