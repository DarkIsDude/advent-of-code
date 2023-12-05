package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ": ")
		card := strings.Replace(split[0], "Card ", "", -1)

		lotery := strings.Split(strings.TrimSpace((split[1])), " | ")

		winning := strings.Split(lotery[0], " ")
		numbers := strings.Split(lotery[1], " ")

		points := 0

		for _, number := range numbers {
			for _, win := range winning {
				fmt.Println(number, win)
				if number == win && number != "" && win != "" {
					if points == 0 {
						points = 1
					} else {
						points *= 2
					}
				}
			}
		}

		sum += points

		fmt.Printf("For card %s, the points are %d\n", card, points)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The total points are %d\n", sum)
}
