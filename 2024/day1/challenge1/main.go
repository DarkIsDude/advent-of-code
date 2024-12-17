package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"math"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbersR := []int{}
	numbersL := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")

		fmt.Println(numbers[3])

		numR, err := strconv.Atoi(numbers[0])
		numL, err := strconv.Atoi(numbers[3])

		if err != nil {
			log.Fatal(err)
		}

		numbersR = append(numbersR, numR)
		numbersL = append(numbersL, numL)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(numbersR)
	sort.Ints(numbersL)

	distance := 0
	for i, _ := range numbersR {
		distance += int(math.Abs(float64(numbersR[i]) - float64(numbersL[i])))
	}

	fmt.Println(distance)
}
