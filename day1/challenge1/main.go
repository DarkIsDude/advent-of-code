package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	countLowerThanPrevious := 0

	scanner.Scan()
	previousNumber, err := strconv.Atoi(scanner.Text())

	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic(err)
		}

		if number > previousNumber {
			countLowerThanPrevious++
		}

		previousNumber = number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(countLowerThanPrevious)
}
