package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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
	foundSymbol := false
	sumOfNumbers := 0

	for i, line := range input {
		for j, char := range line {
			_, err := strconv.Atoi(string(char))

			if err == nil {
				currentNumber += string(char)

				if isSymbolAround(input, i, j) {
					foundSymbol = true
				}
			}

			if err != nil || j == len(line)-1 {
				if currentNumber != "" && foundSymbol {
					number, err := strconv.Atoi(currentNumber)

					if err != nil {
						log.Fatal(err)
					}

					fmt.Println(number)
					sumOfNumbers += number
				}

				currentNumber = ""
				foundSymbol = false
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumOfNumbers)
}

func isSymbolAround(input []string, i int, j int) bool {
	if i-1 > 0 && isSymbolOnLine(input[i-1], j, true) {
		return true
	}

	if isSymbolOnLine(input[i], j, false) {
		return true
	}

	if i+1 < len(input) && isSymbolOnLine(input[i+1], j, true) {
		return true
	}

	return false
}

func isSymbolOnLine(line string, j int, checkCenter bool) bool {
	if j-1 > 0 && isSymbol(line[j-1]) {
		return true
	}

	if checkCenter && isSymbol(line[j]) {
		return true
	}

	if j+1 < len(line) && isSymbol(line[j+1]) {
		return true
	}

	return false
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
