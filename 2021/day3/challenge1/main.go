package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	gammaInitiated := false
	var countNumber []int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Traverse input
	for scanner.Scan() {
		action := scanner.Text()

		for pos, char := range action {
			if !gammaInitiated {
				countNumber = append(countNumber, 0)
			}

			if char == '1' {
				countNumber[pos] += 1
			} else {
				countNumber[pos] -= 1
			}
		}

		gammaInitiated = true
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Create gamma and epsilon
	binaryGamma := ""
	binaryEpsilon := ""

	for _, value := range countNumber {
		if value > 0 {
			binaryGamma += "1"
			binaryEpsilon += "0"
		} else {
			binaryGamma += "0"
			binaryEpsilon += "1"
		}
	}

	fmt.Printf("Binary / gamma : %s / epsilon : %s\n", binaryGamma, binaryEpsilon)

	gamma, err := strconv.ParseInt(binaryGamma, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilon, err := strconv.ParseInt(binaryEpsilon, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result : %d\n", gamma*epsilon)
}
