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
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		numS := -1
		numE := -1

		for i, _ := range line {
			if (numS == -1) {
				num, err := strconv.Atoi(string(line[i]))

				if err == nil {
					numS = num
				}
			}

			if (numE == -1) {
				num, err := strconv.Atoi(string(line[len(line) - 1 - i]))

				if err == nil {
					numE = num
				}
			}
		}

		sum += numS * 10 + numE
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
