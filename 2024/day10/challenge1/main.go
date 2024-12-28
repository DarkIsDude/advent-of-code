package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var expectedValue = map[byte]byte{
	'-': '0',
	'0': '1',
	'1': '2',
	'2': '3',
	'3': '4',
	'4': '5',
	'5': '6',
	'6': '7',
	'7': '8',
	'8': '9',
}

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	topographicMap := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		topographicMap = append(topographicMap, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(findAllPaths(topographicMap))
}

func findAllPaths(topographicMap []string) int {
	sum := 0

	for y, line := range topographicMap {
		for x, char := range line {
			if char == '0' {
				result := emtpyTopographicMap(topographicMap)
				result = findPaths(topographicMap, result, x, y, '-')

				fmt.Println("Path:", x, y)
				displayTopographicMap(result)
				successPaths := countSuccessPaths(result)
				fmt.Println("Success paths:", successPaths)

				sum += successPaths
			}
		}
	}

	return sum
}

func countSuccessPaths(topographicMap []string) int {
	sum := 0

	for _, line := range topographicMap {
		for _, char := range line {
			if char == 'O' {
				sum++
			}
		}
	}

	return sum
}

func emtpyTopographicMap(topographicMap []string) []string {
	empty := []string{}

	for _, line := range topographicMap {
		empty = append(empty, strings.Repeat(".", len(line)))
	}

	return empty
}

func findPaths(topographicMap []string, result []string, startX int, startY int, before byte) []string {
	if startX < 0 || startY < 0 || startX >= len(topographicMap[0]) || startY >= len(topographicMap) {
		return result
	}

	if result[startY][startX] == 'X' {
		return result
	}

	if topographicMap[startY][startX] != expectedValue[before] {
		return result
	}

	if topographicMap[startY][startX] == '9' {
		result[startY] = result[startY][:startX] + "O" + result[startY][startX+1:]
		return result
	}

	result[startY] = result[startY][:startX] + "X" + result[startY][startX+1:]

	result = findPaths(topographicMap, result, startX-1, startY, topographicMap[startY][startX])
	result = findPaths(topographicMap, result, startX+1, startY, topographicMap[startY][startX])
	result = findPaths(topographicMap, result, startX, startY-1, topographicMap[startY][startX])
	result = findPaths(topographicMap, result, startX, startY+1, topographicMap[startY][startX])

	return result
}

func displayTopographicMap(topographicMap []string) {
	for _, line := range topographicMap {
		fmt.Println(line)
	}

	fmt.Println()
}
