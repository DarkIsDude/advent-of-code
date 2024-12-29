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
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	stones := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		stones = strings.Split(line, " ")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	fmt.Println(len(stones))
}

func blink(oldStones []string) []string {
	stones := []string{}

	for _, stone := range oldStones {
		if stone == "0" {
			stones = append(stones, "1")
		} else if (len(stone) % 2) == 0 {
			numberLeft, err := strconv.Atoi(stone[:len(stone)/2])

			if err != nil {
				log.Fatal(err)
			}

			numberRight, err := strconv.Atoi(stone[len(stone)/2:])
			if err != nil {
				log.Fatal(err)
			}

			stones = append(stones, strconv.Itoa(numberLeft))
			stones = append(stones, strconv.Itoa(numberRight))
		} else {
			number, err := strconv.Atoi(stone)

			if err != nil {
				log.Fatal(err)
			}

			stones = append(stones, strconv.Itoa(number*2024))
		}
	}

	fmt.Println(stones)

	return stones
}
