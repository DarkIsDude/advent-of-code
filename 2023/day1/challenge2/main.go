package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var letterToValue map[string]int = map[string]int{
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}


func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		letterS := map[string]int{
			"one": 0,
			"two": 0,
			"three": 0,
			"four": 0,
			"five": 0,
			"six": 0,
			"seven": 0,
			"eight": 0,
			"nine": 0,
		}
		letterE := map[string]int{
			"one": 0,
			"two": 0,
			"three": 0,
			"four": 0,
			"five": 0,
			"six": 0,
			"seven": 0,
			"eight": 0,
			"nine": 0,
		}

		numS := -1
		numE := -1

		for i := range line {
			charS := string(line[i])
			charE := string(line[len(line) - 1 - i])

			if (numS == -1) {
				num, err := strconv.Atoi(charS)

				if err == nil {
					numS = num
				}

				for numberAsString, value := range letterS {
					if charS == string(numberAsString[value]) {
						letterS[numberAsString]++

						if (letterS[numberAsString] == len(numberAsString)) {
							numS = letterToValue[numberAsString]
						}
					} else if charS == string(numberAsString[0]) {
						letterS[numberAsString] = 1
					} else {
						letterS[numberAsString] = 0
					}
				}
			}

			if (numE == -1) {
				num, err := strconv.Atoi(charE)

				if err == nil {
					numE = num
				}

				for numberAsString, value := range letterE {
					if charE == string(numberAsString[len(numberAsString) - 1 - value]) {
						letterE[numberAsString]++

						if (letterE[numberAsString] == len(numberAsString)) {
							numE = letterToValue[numberAsString]
						}
					} else if charE == string(numberAsString[len(numberAsString) - 1]) {
						letterE[numberAsString] = 1
					} else {
						letterE[numberAsString] = 0
					}
				}
			}
		}

		fmt.Println(numS, numE)
		sum += numS * 10 
		fmt.Println(sum)
		sum += numE
		fmt.Println(sum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
