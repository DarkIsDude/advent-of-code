package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// 0 use 6 segments
// 1 use 2 segments OK
// 2 use 5 segments
// 3 use 5 segments
// 4 use 4 segments OK
// 5 use 5 segments
// 6 use 6 segments
// 7 use 3 segments OK
// 8 use 7 segments OK
// 9 use 6 segments

type Display struct {
	signals []string
	digits  []string
}

func main() {
	displays := readFile()

	oneAppear := 0
	fourAppear := 0
	sevenAppear := 0
	heightAppear := 0

	for _, display := range displays {
		for _, digit := range display.digits {
			switch len(digit) {
			case 2:
				oneAppear++
			case 4:
				fourAppear++
			case 3:
				sevenAppear++
			case 7:
				heightAppear++
			}
		}
	}

	fmt.Printf("1 appear %d\n", oneAppear)
	fmt.Printf("4 appear %d\n", fourAppear)
	fmt.Printf("7 appear %d\n", sevenAppear)
	fmt.Printf("8 appear %d\n", heightAppear)

	fmt.Printf("Total %d\n", oneAppear+fourAppear+sevenAppear+heightAppear)
}

func readFile() []Display {
	var displays []Display

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if splittedLine := strings.Split(line, " | "); len(splittedLine) == 2 {
			displays = append(displays, Display{
				signals: strings.Split(splittedLine[0], " "),
				digits:  strings.Split(splittedLine[1], " "),
			})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return displays
}
