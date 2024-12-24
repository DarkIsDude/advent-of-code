package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	result   int
	elements []int

	currentResult int
}

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		eq := equation{}
		result, err := strconv.Atoi(strings.Split(line, ": ")[0])

		if err != nil {
			log.Fatal(err)
		}

		eq.result = result
		elementStrings := strings.Split(strings.Split(line, ": ")[1], " ")

		for _, elementString := range elementStrings {
			element, err := strconv.Atoi(elementString)

			if err != nil {
				log.Fatal(err)
			}

			eq.elements = append(eq.elements, element)
		}

		eq.currentResult = eq.elements[0]
		eq.elements = eq.elements[1:]

		fmt.Printf("Resolve %d (%d): %v\n", eq.result, eq.currentResult, eq.elements)

		if canResolve(eq) {
			sum += eq.result
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum:", sum)
}

func canResolve(eq equation) bool {
	if len(eq.elements) < 1 {
		return eq.currentResult == eq.result
	}

	// With addition
	canWithAddition := canResolve(equation{
		result:        eq.result,
		elements:      eq.elements[1:],
		currentResult: eq.currentResult + eq.elements[0],
	})

	// With multiplication
	canWithMultiplication := canResolve(equation{
		result:        eq.result,
		elements:      eq.elements[1:],
		currentResult: eq.currentResult * eq.elements[0],
	})

	// With concatenation
	canWithConcatenation := canResolve(equation{
		result:        eq.result,
		elements:      eq.elements[1:],
		currentResult: concat(eq.currentResult, eq.elements[0]),
	})

	return canWithAddition || canWithMultiplication || canWithConcatenation
}

func concat(a, b int) int {
	n, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))

	if err != nil {
		log.Fatal(err)
	}

	return n
}
