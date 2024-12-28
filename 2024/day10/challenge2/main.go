package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
				paths := findPaths(topographicMap, [][]int{}, x, y, '-')

				fmt.Println("Path:", x, y)
				fmt.Println("Success paths:", paths)

				sum += paths
			}
		}
	}

	return sum
}

func findPaths(topographicMap []string, path [][]int, startX int, startY int, before byte) int {
	path = append(path, []int{startX, startY})

	if startX < 0 || startY < 0 || startX >= len(topographicMap[0]) || startY >= len(topographicMap) {
		return 0
	}

	if topographicMap[startY][startX] != expectedValue[before] {
		return 0
	}

	if topographicMap[startY][startX] == '9' {
		return 1
	}

	return findPaths(topographicMap, path, startX-1, startY, topographicMap[startY][startX]) +
		findPaths(topographicMap, path, startX+1, startY, topographicMap[startY][startX]) +
		findPaths(topographicMap, path, startX, startY-1, topographicMap[startY][startX]) +
		findPaths(topographicMap, path, startX, startY+1, topographicMap[startY][startX])
}
