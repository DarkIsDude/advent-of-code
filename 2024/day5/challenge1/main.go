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

			if sectionMakeSense(numbers, mustBeBefore) {
				fmt.Println("Match", line)
				sum += numbers[len(numbers)/2]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum", sum)
}

func sectionMakeSense(numbers []int, mustBeBefore map[int][]int) bool {
	for i := len(numbers) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if slices.Contains(mustBeBefore[numbers[i]], numbers[j]) {
				fmt.Println("Doesn't make sense", numbers[i], numbers[j])
				return false
			}
		}
	}

	return true
}
