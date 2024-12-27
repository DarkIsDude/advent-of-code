package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Space struct {
	used   bool
	index  int
	length int
}

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		for _, char := range line {
			number, err := strconv.Atoi(string(char))

			if err != nil {
				log.Fatal(err)
			}

			numbers = append(numbers, number)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	diskMap := expandDiskMap(numbers)
	displayDiskMap(diskMap)
	diskMap = compactSmart(diskMap)
	displayDiskMap(diskMap)
	fmt.Println(checksum(diskMap))
}

func checksum(diskMap []Space) int {
	checksum := 0
	index := 0

	for _, space := range diskMap {
		for i := 0; i < space.length; i++ {
			if space.used {
				checksum += index * space.index
			}
			index++
		}
	}

	return checksum
}

func compactSmart(diskMap []Space) []Space {
	currentIndex := diskMap[len(diskMap)-1].index

	for currentIndex > 0 {
		fmt.Println("Current index:", currentIndex)

		for i, space := range diskMap {
			if space.used && space.index == currentIndex {
				nextFreeSpaceIndex := getIndexOfNextFreeSpace(diskMap, 0, space.length)
				fmt.Println("Next free space index:", nextFreeSpaceIndex)

				if nextFreeSpaceIndex >= 0 {
					oldLength := space.length
					lengthDiff := int(math.Abs(float64(oldLength - diskMap[nextFreeSpaceIndex].length)))

					diskMap[nextFreeSpaceIndex].used = true
					diskMap[nextFreeSpaceIndex].index = currentIndex
					diskMap[nextFreeSpaceIndex].length = space.length
					diskMap[i].used = false
					diskMap[i].index = 0
					diskMap[i].length = oldLength

					fmt.Println("Old length:", oldLength)
					fmt.Println("Length diff:", lengthDiff)

					if lengthDiff > 0 {
						diskMap = slices.Insert(diskMap, nextFreeSpaceIndex+1, Space{used: false, index: 0, length: lengthDiff})
					}
				}
			}
		}

		currentIndex--
	}

	return diskMap
}

func getIndexOfNextFreeSpace(diskMap []Space, index int, minLenght int) int {
	for i := index; i < len(diskMap); i++ {
		if diskMap[i].used == false && diskMap[i].length >= minLenght {
			return i
		}
	}

	return -1
}

func displayDiskMap(diskMap []Space) {
	for _, space := range diskMap {
		if space.used {
			fmt.Print(strings.Repeat(strconv.Itoa(space.index), space.length))
		} else {
			fmt.Print(strings.Repeat(".", space.length))
		}
	}

	fmt.Println()
}

func expandDiskMap(diskMap []int) []Space {
	expanded := []Space{}
	index := 0
	freeSpace := false

	for _, space := range diskMap {
		if freeSpace {
			expanded = append(expanded, Space{used: false, length: space})
		} else {
			expanded = append(expanded, Space{used: true, index: index, length: space})
			index++
		}

		freeSpace = !freeSpace
	}

	return expanded
}
