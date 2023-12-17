package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type combinationType struct {
	index int
	indications []int
	final string
}

var cache = map[string]int{}

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumOfArrangements := 0

	for scanner.Scan() {
		line := scanner.Text()

		lineSplited := strings.Split(line, " ")
		spring := lineSplited[0]
	  indicationsAsString := lineSplited[1]

		indications := []int{}
		for _, indication := range strings.Split(indicationsAsString, ",") {
			i, err := strconv.Atoi(indication)

			if err != nil {
				log.Fatal(err)
			}

			indications = append(indications, i)
		}

		spring = spring + "?" + spring + "?" + spring + "?" + spring + "?" + spring
		indications5 := append(indications, indications...)
		indications5 = append(indications5, indications...)
		indications5 = append(indications5, indications...)
		indications5 = append(indications5, indications...)

		fmt.Printf("Spring: %s, %v\n", spring, indications5)

		number := numberOfArrangements(spring, indications5)
		sumOfArrangements += number
		fmt.Println(number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumOfArrangements)
}

func numberOfArrangements(spring string, indications []int) int {
	fromCache := getFromCache(spring, indications)
	if fromCache != -1 {
		return fromCache
	}

	if len(indications) == 0 {
		if strings.IndexByte(spring, '#') == -1 {
			return saveInCache(spring, indications, 1)
		}

		return saveInCache(spring, indications, 0)
	}

	if len(spring) == 0 {
		return saveInCache(spring, indications, 0)
	}

	indication := indications[0]
	_numberOfArrangements := 0

	if spring[0] != '#' {
		_numberOfArrangements += numberOfArrangements(spring[1:], indications)
	}

	if len(spring) < indication {
		return saveInCache(spring, indications, _numberOfArrangements)
	}

	if len(indications) > 1 {
		if len(spring) < indication + 1 {
			return saveInCache(spring, indications, _numberOfArrangements)
		}

		if spring[indication] == '#' {
			return saveInCache(spring, indications, _numberOfArrangements)
		}
	}

	valueExtracted := spring[:indication]
	if strings.IndexByte(valueExtracted, '.') != -1 {
		return saveInCache(spring, indications, _numberOfArrangements)
	}

	value := 0
	if len(indications) > 1 {
		value = _numberOfArrangements + numberOfArrangements(spring[indication + 1:], indications[1:])
	} else {
		value = _numberOfArrangements + numberOfArrangements(spring[indication:], indications[1:])
	}

	return saveInCache(spring, indications, value)
}

func saveInCache(spring string, indications []int, value int) int {
	cache[fmt.Sprintf("%s-%v", spring, indications)] = value

	return value
}

func getFromCache(spring string, indications []int) int {
	value, ok := cache[fmt.Sprintf("%s-%v", spring, indications)]

	if !ok {
		return -1
	}

	return value
}
