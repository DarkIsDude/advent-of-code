package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Display struct {
	signals []string
	digits  []string
}

func main() {
	displays := readFile()
	sum := 0

	for _, display := range displays {
		value := findDigit(display)

		fmt.Println(value)

		sum += value
	}

	fmt.Printf("Final %d\n", sum)
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

func findDigit(display Display) int {
	var signalsFive []string
	var signalsSix []string
	signals := map[int]string{
		0: "",
		1: "",
		2: "",
		3: "",
		4: "",
		5: "",
		6: "",
		7: "",
		8: "",
		9: "",
	}

	for _, signal := range display.signals {
		switch len(signal) {
		case 2:
			signals[1] = signal
		case 4:
			signals[4] = signal
		case 3:
			signals[7] = signal
		case 7:
			signals[8] = signal
		case 5:
			signalsFive = append(signalsFive, signal)
		case 6:
			signalsSix = append(signalsSix, signal)
		default:
			panic("Should never happen this len")
		}
	}

	// At this 1, 4, 7 and 8 are defined

	// 3 : If I remove 1 to all, the one with 3 remaining
	// 5 : Then, if I remove 4, the one with 2 remaining
	// 2 : The last one
	for _, signal := range signalsFive {
		minusOne := removeCommonCharacter(signal, signals[1])
		if len(minusOne) == 3 {
			signals[3] = signal
			continue
		}

		minusFour := removeCommonCharacter(signal, signals[4])
		if len(minusFour) == 2 {
			signals[5] = signal
			continue
		}

		signals[2] = signal
	}

	// At this 1, 2, 3, 4, 5, 7 and 8 are defined

	// 6 : If I remove 1, the one with 5 remaining
	// 9 : If I remove 3, the one with 1 remaining
	// 0 : The last one
	for _, signal := range signalsSix {
		minusOne := removeCommonCharacter(signal, signals[1])
		if len(minusOne) == 5 {
			signals[6] = signal
			continue
		}

		minusThree := removeCommonCharacter(minusOne, signals[3])
		if len(minusThree) == 1 {
			signals[9] = signal
			continue
		}

		signals[0] = signal
	}

	finalResult := 0
	for position, digit := range display.digits {
		value := 0

		for number, signal := range signals {
			diff := removeCommonCharacter(signal, digit)
			if len(signal) == len(digit) && len(diff) == 0 {
				value = number
			}
		}

		multiplier := 1000
		for i := 0; i < position; i++ {
			multiplier = multiplier / 10
		}

		finalResult += value * multiplier
	}

	return finalResult
}

func removeCommonCharacter(toUpdate string, source string) string {
	substracted := ""

	for _, character := range toUpdate {
		present := false

		for _, c := range source {
			if character == c {
				present = true
			}
		}

		if !present {
			substracted += string(character)
		}
	}

	return substracted
}
