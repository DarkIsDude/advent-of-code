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
	stones := map[string]int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		stonesL := strings.Split(line, " ")
		for _, stone := range stonesL {
			stones[stone] = 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 75; i++ {
		fmt.Println(len(stones))
		stones = blink(stones)
	}

	sum := 0
	for _, count := range stones {
		sum += count
	}

	fmt.Println(sum)
}

func blink(oldStones map[string]int) map[string]int {
	stones := map[string]int{}

	for marker, count := range oldStones {
		if marker == "0" {
			stones["1"] += count
		} else if (len(marker) % 2) == 0 {
			numberLeft, err := strconv.Atoi(marker[:len(marker)/2])

			if err != nil {
				log.Fatal(err)
			}

			numberRight, err := strconv.Atoi(marker[len(marker)/2:])
			if err != nil {
				log.Fatal(err)
			}

			stones[strconv.Itoa(numberLeft)] += count
			stones[strconv.Itoa(numberRight)] += count
		} else {
			number, err := strconv.Atoi(marker)

			if err != nil {
				log.Fatal(err)
			}

			stones[strconv.Itoa(number*2024)] += count
		}
	}

	return stones
}
