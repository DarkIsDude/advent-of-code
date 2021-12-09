package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readFile()

	maxX := len(numbers[0]) - 1
	maxY := len(numbers) - 1

	var lowestNumbers []int

	for y, line := range numbers {
		for x, number := range line {
			lowest := true

			// Left
			if x-1 >= 0 && number >= numbers[y][x-1] {
				lowest = false
			}

			// Top
			if y-1 >= 0 && number >= numbers[y-1][x] {
				lowest = false
			}

			// Right
			if x+1 <= maxX && number >= numbers[y][x+1] {
				lowest = false
			}

			// Bottom
			if y+1 <= maxY && number >= numbers[y+1][x] {
				lowest = false
			}

			if lowest {
				lowestNumbers = append(lowestNumbers, number)
			}
		}
	}

	sum := 0
	for _, number := range lowestNumbers {
		sum += number + 1
	}

	fmt.Printf("Numbers found %v\n", lowestNumbers)
	fmt.Printf("Sum %v\n", sum)
}

func readFile() [][]int {
	var numbers [][]int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers = append(numbers, []int{})

		for _, character := range line {
			integer, err := strconv.Atoi(string(character))
			if err != nil {
				panic(err)
			}

			numbers[lineNumber] = append(numbers[lineNumber], integer)
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return numbers
}
