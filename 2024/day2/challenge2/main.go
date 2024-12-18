package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	countValid := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbersS := strings.Split(line, " ")
		numbers := make([]int, len(numbersS))

		for i, numberS := range numbersS {
			number, _ := strconv.Atoi(numberS)
			numbers[i] = number
		}

		validLine := isLineSafeWithOneJoker(numbers)

		if validLine {
			countValid++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(countValid)
}

func isLineSafeWithOneJoker(numbers []int) bool {
	validLine := isLineSafe(numbers)

	if validLine {
		return true
	}

	for i := 0; i < len(numbers); i++ {
		newNumbers := make([]int, len(numbers))
		copy(newNumbers, numbers)
		newNumbers = append(newNumbers[:i], newNumbers[i+1:]...)

		validLine = isLineSafe(newNumbers)

		if validLine {
			return true
		}
	}

	return false
}

func isLineSafe(numbers []int) bool {
	var direction int

	if numbers[0] > numbers[1] {
		direction = -1
	} else {
		direction = 1
	}

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] && direction == -1 {
			return false
		}

		if numbers[i] < numbers[i-1] && direction == 1 {
			return false
		}

		distance := int(math.Abs(float64(numbers[i] - numbers[i-1])))
		if distance < 1 || distance > 3 {
			return false
		}
	}

	return true
}
