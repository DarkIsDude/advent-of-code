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
	file, err := os.Open("./input.txt")

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

		var direction int

		if numbers[0] > numbers[1] {
			direction = -1
		} else {
			direction = 1
		}

		validLine := true
		for i := 1; i < len(numbers); i++ {
			if numbers[i] > numbers[i-1] && direction == -1 {
				validLine = false
				break
			}

			if numbers[i] < numbers[i-1] && direction == 1 {
				validLine = false
				break
			}

			distance := int(math.Abs(float64(numbers[i] - numbers[i-1])))
			if distance < 1 || distance > 3 {
				validLine = false
				break
			}
		}

		if validLine {
			countValid++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(countValid)
}
