package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const FILE_LOCATION = "./input.txt"
const ITERATION = 40
const DEBUG = false

func main() {
	path, combinaisons := readFile()
	counter := make(map[string]int) // XX => value

	for i := 0; i < len(path)-1; i++ {
		counter[fmt.Sprintf("%s%s", string(path[i]), string(path[i+1]))]++
	}

	for i := 0; i < ITERATION; i++ {
		fmt.Printf("Iteration %d\n", i+1)

		newCounter := make(map[string]int)
		for combinaison, toAdd := range combinaisons {
			suffix := fmt.Sprintf("%s%s", string(combinaison[0]), toAdd)
			prefix := fmt.Sprintf("%s%s", toAdd, string(combinaison[1]))

			if DEBUG {
				fmt.Printf("The value of %s and %s will be %d (%s)\n", suffix, prefix, counter[combinaison], combinaison)
			}

			newCounter[suffix] += counter[combinaison]
			newCounter[prefix] += counter[combinaison]
		}

		counter = newCounter
	}

	for combinaison, value := range counter {
		fmt.Printf("We have %d %s\n", value, combinaison)
	}

	fmt.Println("-------")

	singleCounter := computeToSingleValue(counter)

	for combinaison, value := range singleCounter {
		fmt.Printf("We have %d %s\n", value, combinaison)
	}

	min, max := extractMinMax(singleCounter)

	fmt.Println(max - min)
}

func readFile() (string, map[string]string) {
	combinaisons := make(map[string]string)

	file, err := os.Open(FILE_LOCATION)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	path := scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		combinaison := strings.Split(scanner.Text(), " -> ")
		combinaisons[combinaison[0]] = combinaison[1]
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return path, combinaisons
}

func computeToSingleValue(counter map[string]int) map[string]int {
	singleCounter := make(map[string]int)

	for combinaison, value := range counter {
		singleCounter[string(combinaison[0])] += value
		singleCounter[string(combinaison[1])] += value
	}

	for c, value := range singleCounter {
		singleCounter[c] = value/2 + value%2
	}

	return singleCounter
}

func extractMinMax(counter map[string]int) (int, int) {
	max := -1
	min := -1

	for _, value := range counter {
		if max == -1 {
			max = value
			min = value
		}

		if value > max {
			max = value
		}

		if value < min {
			min = value
		}
	}

	return min, max
}
