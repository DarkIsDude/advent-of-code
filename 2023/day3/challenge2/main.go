package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type coord struct {
	i int
	j int
}

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
		input = append(input, line)
	}

	currentNumber := ""
	foundSymbol := map[coord]bool{}
	symbolToNumber := map[coord][]int{}

	for i, line := range input {
		for j, char := range line {
			_, err := strconv.Atoi(string(char))

			if err == nil {
				currentNumber += string(char)

				for _, position := range symbolAround(input, i, j) {
					foundSymbol[position] = true
				}
			}

			if err != nil || j == len(line)-1 {
				if currentNumber != "" {
					number, err := strconv.Atoi(currentNumber)

					if err != nil {
						log.Fatal(err)
					}

					for symbol := range foundSymbol {
						symbolToNumber[symbol] = append(symbolToNumber[symbol], number)
					}
				}

				currentNumber = ""
				foundSymbol = map[coord]bool{}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, numbers := range symbolToNumber {
		if len(numbers) == 2 {
			sum += numbers[0] * numbers[1]
		}
	}

	fmt.Println(sum)
}

func symbolAround(input []string, i int, j int) []coord {
	symbolPosition := []coord{}
	if i-1 > 0 {
		symbolOnLine := symbolOnLine(input[i-1], j, true)
		for _, position := range symbolOnLine {
			symbolPosition = append(symbolPosition, coord{i - 1, position})
		}
	}

	if i+1 < len(input) {
		symbolOnLine := symbolOnLine(input[i+1], j, true)
		for _, position := range symbolOnLine {
			symbolPosition = append(symbolPosition, coord{i + 1, position})
		}
	}

	symbolOnLine := symbolOnLine(input[i], j, false)
	for _, position := range symbolOnLine {
		symbolPosition = append(symbolPosition, coord{i, position})
	}

	return symbolPosition
}

func symbolOnLine(line string, j int, checkCenter bool) []int {
	symbolPosition := []int{}

	if j-1 > 0 && isSymbol(line[j-1]) {
		symbolPosition = append(symbolPosition, j-1)
	}

	if checkCenter && isSymbol(line[j]) {
		symbolPosition = append(symbolPosition, j)
	}

	if j+1 < len(line) && isSymbol(line[j+1]) {
		symbolPosition = append(symbolPosition, j+1)
	}

	return symbolPosition
}

func isSymbol(char byte) bool {
	if char == '.' {
		return false
	}

	_, err := strconv.Atoi(string(char))

	if err != nil {
		return true
	}

	return false
}
