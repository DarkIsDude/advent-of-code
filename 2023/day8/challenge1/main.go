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
	scanner.Scan()
	path := strings.TrimSpace(scanner.Text())
	scanner.Scan()
	nodes := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		nodes[line[0:3]] = []string{line[7:10], line[12:15]}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	currentNode := "AAA"
	currentPathIndex := 0
	numberOfIterations := 0
	for currentNode != "ZZZ" {
		if path[currentPathIndex] == 'L' {
			currentNode = nodes[currentNode][0]
		} else {
			currentNode = nodes[currentNode][1]
		}

		numberOfIterations++
		currentPathIndex = (currentPathIndex + 1) % len(path)
	}

	fmt.Println(numberOfIterations)
}
