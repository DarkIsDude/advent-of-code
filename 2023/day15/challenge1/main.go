package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		sequence := strings.Split(line, ",")
		sum := 0

		for _, value := range sequence {
			sum += runHASHAlgotithm(value)
		}

		fmt.Println(sum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func runHASHAlgotithm(line string) int {
	hash := 0

	for _, char := range line {
		hash += int(char)
		hash *= 17
		hash = hash % 256
	}

	return hash
}
