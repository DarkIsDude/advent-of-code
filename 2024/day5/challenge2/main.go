package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	mustBeBefore := make(map[int][]int)
	sum := 0

	scanRules := true

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Reading", line)

		if line == "" {
			scanRules = false
			fmt.Println("Switching to scan numbers", mustBeBefore)
			continue
		}

		if scanRules {
			numbers := strings.Split(line, "|")
			number1, _ := strconv.Atoi(numbers[0])
			number2, _ := strconv.Atoi(numbers[1])

			mustBeBefore[number1] = append(mustBeBefore[number1], number2)
		} else {
			numbersS := strings.Split(line, ",")
			numbers := make([]int, 0)

			for _, numberS := range numbersS {
				number, _ := strconv.Atoi(numberS)
				numbers = append(numbers, number)
			}

			ok, _, _ := sectionMakeSense(numbers, mustBeBefore)

			if !ok {
				fmt.Println("Match", line)
				fixedSection := fixSection(numbers, mustBeBefore)
				sum += fixedSection[len(fixedSection)/2]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum", sum)
}

func fixSection(numbers []int, mustBeBefore map[int][]int) []int {
	section := make([]int, len(numbers))
	copy(section, numbers)

	makeSense := false

	for !makeSense {
		makeSenseL, i, j := sectionMakeSense(section, mustBeBefore)

		if !makeSenseL {
			section[i], section[j] = section[j], section[i]
		} else {
			makeSense = true
		}
	}

	return section
}

func sectionMakeSense(numbers []int, mustBeBefore map[int][]int) (bool, int, int) {
	for i := len(numbers) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if slices.Contains(mustBeBefore[numbers[i]], numbers[j]) {
				fmt.Println("Doesn't make sense", numbers[i], numbers[j])
				return false, i, j
			}
		}
	}

	return true, -1, -1
}
