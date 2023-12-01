package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const FILE_LOCATION = "./input.txt"
const ITERATION = 10
const DEBUG = false

func main() {
	path, combinaisons := readFile()

	for ite := 0; ite < ITERATION; ite++ {
		fmt.Printf("Will start iteration %d\n", ite+1)
		if DEBUG {
			fmt.Println(path)
		}

		for i := len(path) - 2; i >= 0; i-- {
			stringToTest := fmt.Sprintf("%s%s", string(path[i]), string(path[i+1]))

			if DEBUG {
				fmt.Printf("Checking at position %d : %s\n", i, stringToTest)
			}

			if val, ok := combinaisons[stringToTest]; ok {
				if DEBUG {
					fmt.Printf("Adding %s to %s\n", val, path)
				}

				path = fmt.Sprintf("%s%s%s", path[:i+1], val, path[i+1:])

				if DEBUG {
					fmt.Printf("Result %s to %s\n", val, path)
				}
			}
		}
	}

	if DEBUG {
		fmt.Printf("The final result after %d iteration is %s\n", ITERATION, path)
	}

	counter := countEachCaracter(path)
	min, max := extractMinMax(counter)

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

func countEachCaracter(s string) map[rune]int {
	counter := make(map[rune]int)

	for _, c := range s {
		counter[c]++
	}

	return counter
}

func extractMinMax(counter map[rune]int) (int, int) {
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
