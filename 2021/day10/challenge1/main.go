package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const DEBUG = false

func main() {
	lines := readFile()
	points := 0

	for _, line := range lines {
		valid, character := processLine(line)

		if !valid {
			switch character {
			case ')':
				points += 3
			case ']':
				points += 57
			case '}':
				points += 1197
			case '>':
				points += 25137
			}
		}
	}

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

func processLine(line string) (bool, rune) {
	var openers []rune
	valid := true
	wrongValue := ' '

	if DEBUG {
		fmt.Printf("Starting %s\n", line)
	}

	for _, c := range line {
		if !valid {
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
				valid = false
				wrongValue = c
			} else {
				openers = openers[:len(openers)-1]
			}
		}
	}

	fmt.Printf("%s is valid ? %t (%s)\n", line, valid, string(wrongValue))

	return valid, wrongValue
}
