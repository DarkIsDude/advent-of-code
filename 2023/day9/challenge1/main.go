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
	sumOfGuessValues := 0

	for scanner.Scan() {
		line := scanner.Text()

		valuesAsString := strings.Split(line, " ")
		values := make([]int, len(valuesAsString))

		for i, valueAsString := range valuesAsString {
			n, err := strconv.Atoi(valueAsString)

			if err != nil {
				log.Fatal(err)
			}

			values[i] = n
		}

		guessValue := guessNextValue(values)
		sumOfGuessValues += guessValue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumOfGuessValues)
}

func guessNextValue(values []int) int {
	fmt.Println(values)

	if len(values) == 0 {
		log.Fatal("No values")
	}

	allValuesEqualZero := true

	for _, value := range values {
		if value != 0 {
			allValuesEqualZero = false
			break
		}
	}

	if allValuesEqualZero {
		return 0
	}

	differences := []int{}

	for i := 0; i < len(values)-1; i++ {
		differences = append(differences, values[i+1]-values[i])
	}

	nextValueOfDifferences := guessNextValue(differences)

	return values[len(values)-1] + nextValueOfDifferences
}
