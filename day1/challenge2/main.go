package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	countLowerThanPrevious := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	previousPreviousPreviousNumber := -1
	previousPreviousNumber := -1
	previousNumber := -1

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if previousPreviousPreviousNumber < 0 {
			previousPreviousPreviousNumber = number
			continue
		}

		if previousPreviousNumber < 0 {
			previousPreviousNumber = number
			continue
		}

		if previousNumber < 0 {
			previousNumber = number
			continue
		}

		currentNumber := number

		oldSum := previousNumber + previousPreviousNumber + previousPreviousPreviousNumber
		newSum := currentNumber + previousNumber + previousPreviousNumber

		fmt.Println(previousNumber, previousPreviousNumber, previousPreviousPreviousNumber)
		fmt.Println(currentNumber, previousNumber, previousPreviousNumber)
		fmt.Println(oldSum, newSum)
		fmt.Println("-------")
		if newSum > oldSum {
			countLowerThanPrevious++
		}

		previousPreviousPreviousNumber = previousPreviousNumber
		previousPreviousNumber = previousNumber
		previousNumber = currentNumber
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(countLowerThanPrevious)
}
