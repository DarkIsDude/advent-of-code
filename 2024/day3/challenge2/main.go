package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	enabled := true

	for scanner.Scan() {
		line := scanner.Text()

		reg := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
		matches := reg.FindAllString(line, -1)

		for _, match := range matches {
			if match == "do()" {
				enabled = true
			} else if match == "don't()" {
				enabled = false
			} else if enabled {
				reg = regexp.MustCompile(`\d+`)
				numbers := reg.FindAllString(match, -1)

				num1, _ := strconv.Atoi(numbers[0])
				num2, _ := strconv.Atoi(numbers[1])

				sum += num1 * num2
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
