package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Space struct {
	used  bool
	index int
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
	diskMap = compact(diskMap)
	displayDiskMap(diskMap)
	fmt.Println(checksum(diskMap))
}

func checksum(diskMap []Space) int {
	checksum := 0

	for index, space := range diskMap {
		if space.used {
			checksum += index * space.index
		}
	}

	return checksum
}

func compact(diskMap []Space) []Space {
	indexOfFreeSpace := getIndexOfNextFreeSpace(diskMap, 0)
	indexToCompute := len(diskMap) - 1

	for indexOfFreeSpace < indexToCompute {
		if diskMap[indexToCompute].used {
			diskMap[indexOfFreeSpace].used = true
			diskMap[indexOfFreeSpace].index = diskMap[indexToCompute].index
			diskMap[indexToCompute].used = false
			indexOfFreeSpace = getIndexOfNextFreeSpace(diskMap, indexOfFreeSpace)
		}

		indexToCompute--
	}

	return diskMap
}

func getIndexOfNextFreeSpace(diskMap []Space, index int) int {
	for i := index; i < len(diskMap); i++ {
		if diskMap[i].used == false {
			return i
		}
	}

	log.Fatal("No free space found")
	return -1
}

func displayDiskMap(diskMap []Space) {
	for _, space := range diskMap {
		if space.used {
			fmt.Print(strconv.Itoa(space.index))
		} else {
			fmt.Print(".")
		}
	}

	fmt.Println()
}

func expandDiskMap(diskMap []int) []Space {
	expanded := []Space{}
	index := 0
	freeSpace := false

	for _, space := range diskMap {
		for i := 0; i < space; i++ {
			if freeSpace {
				expanded = append(expanded, Space{used: false})
			} else {
				expanded = append(expanded, Space{used: true, index: index})
			}
		}

		if freeSpace {
			index++
		}

		freeSpace = !freeSpace
	}

	return expanded
}
