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

	currentNodes := []string{}
	for node := range nodes {
		if node[2] == 'A' {
			currentNodes = append(currentNodes, node)
		}
	}

	fmt.Println(currentNodes)

	nodesToStep := []int{}
	currentSteps := []int{}
	for i, node := range currentNodes {
		nodesToStep = append(nodesToStep, calculateStep(node, nodes, path))
		currentSteps = append(currentSteps, nodesToStep[i])

		fmt.Printf("%s: %d\n", node, calculateStep(node, nodes, path))
	}

	for !allEqual(currentSteps) {
		lowestIndex := lowestIndex(currentSteps)
		currentSteps[lowestIndex] = currentSteps[lowestIndex] + nodesToStep[lowestIndex]
	}

	fmt.Println(currentSteps[0])
}

func lowestIndex(numbers []int) int {
	lowestIndex := 0

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[lowestIndex] {
			lowestIndex = i
		}
	}

	return lowestIndex
}

func allEqual(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != numbers[0] {
			return false
		}
	}

	return true
}

func allPathEndsWithZ(paths []string) bool {
	for _, path := range paths {
		if path[2] != 'Z' {
			return false
		}
	}

	return true
}

func calculateStep(start string, nodes map[string][]string, path string) int {
	currentNode := start
	currentPathIndex := 0
	numberOfIterations := 0

	for currentNode[2] != 'Z' {
		if path[currentPathIndex] == 'L' {
			currentNode = nodes[currentNode][0]
		} else {
			currentNode = nodes[currentNode][1]
		}

		numberOfIterations++
		currentPathIndex = (currentPathIndex + 1) % len(path)
	}

	return numberOfIterations
}
