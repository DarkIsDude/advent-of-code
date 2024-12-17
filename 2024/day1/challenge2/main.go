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
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbersL := []int{}
	numbersR := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		numL, err := strconv.Atoi(numbers[0])
		numR, err := strconv.Atoi(numbers[3])

		if err != nil {
			log.Fatal(err)
		}

		numbersL = append(numbersL, numL)
		numbersR = append(numbersR, numR)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	appearances := make(map[int]int)
	for i := 0; i < len(numbersR); i++ {
		appearances[numbersR[i]]++
	}

	distance := 0
	for i := 0; i < len(numbersL); i++ {
		fmt.Println(numbersL[i], appearances[numbersL[i]])
		distance += numbersL[i] * appearances[numbersL[i]]
	}

	fmt.Println(distance)
}
