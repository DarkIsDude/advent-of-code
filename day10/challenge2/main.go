package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const DEBUG = false

func main() {
	lines := readFile()
	points := 0
	var validPoint []int

	for _, line := range lines {
		valid, point := processLine(line)

		if valid {
			validPoint = append(validPoint, point)
		}
	}

	sort.Ints(validPoint)
	points += validPoint[len(validPoint)/2]

	fmt.Printf("Points : %d\n", points)
}

func readFile() []string {
	var lines []string

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func processLine(line string) (bool, int) {
	points := 0
	var openers []rune
	wrongValue := ' '

	if DEBUG {
		fmt.Printf("Starting %s\n", line)
	}

	for _, c := range line {
		if wrongValue != ' ' {
			continue
		}

		openerExpected := ' '

		if DEBUG {
			fmt.Printf("Processing %s\n", string(c))
		}

		switch c {
		case '<', '[', '{', '(':
			openers = append(openers, c)

			if DEBUG {
				fmt.Printf("Opener detected %s %s\n", string(c), string(openers))
			}
		case ')':
			if DEBUG {
				fmt.Println(") detected, expecting (")
			}

			openerExpected = '('
		case '}':
			if DEBUG {
				fmt.Println("} detected, expecting {")
			}

			openerExpected = '{'
		case ']':
			if DEBUG {
				fmt.Println("] detected, expecting [")
			}

			openerExpected = '['
		case '>':
			if DEBUG {
				fmt.Println("> detected, expecting <")
			}

			openerExpected = '<'
		default:
			panic("Should never happen")
		}

		if openerExpected != ' ' {
			lastOpeners := openers[len(openers)-1]

			if lastOpeners != openerExpected {
				wrongValue = c
			} else {
				openers = openers[:len(openers)-1]
			}
		}
	}

	if wrongValue != ' ' {
		switch wrongValue {
		case ')':
			points += 3
		case ']':
			points += 57
		case '}':
			points += 1197
		case '>':
			points += 25137
		}

		if DEBUG {
			fmt.Printf("%s is not valid because of %s and points %d\n", line, string(wrongValue), points)
		}
	} else {
		var closers []rune
		for i := len(openers) - 1; i >= 0; i-- {
			closers = append(closers, openers[i])
		}

		for _, closer := range closers {
			points *= 5

			switch closer {
			case '(':
				points += 1
			case '{':
				points += 3
			case '[':
				points += 2
			case '<':
				points += 4
			}
		}

		fmt.Printf("%s is valid but missing %s with points %d\n", line, string(openers), points)
	}

	return wrongValue == ' ', points
}
