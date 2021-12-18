package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
)

const FILE_LOCATION = "./input.txt"
const DEBUG = false

func main() {
	maximumMagnitude := 0
	snailfishs := readFile()

	for i := 0; i < len(snailfishs); i++ {
		for j := 0; j < len(snailfishs); j++ {
			snailfishs := readFile()

			if i == j {
				continue
			}

			snailfishIJAsArray := []interface{}{reduce(snailfishs[i]), reduce(snailfishs[j])}
			snailfishIJ := reduce(interface{}(snailfishIJAsArray))
			snailfishIJAsArray = snailfishIJ.([]interface{})
			magnitudeIJ := magnitude(snailfishIJAsArray)

			if magnitudeIJ > maximumMagnitude {
				maximumMagnitude = magnitudeIJ
			}
		}
	}

	fmt.Println(maximumMagnitude)
}

// ### READING FILE

func readFile() []interface{} {
	var snailfishs []interface{}

	file, err := os.Open(FILE_LOCATION)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		snailfish, _ := readSnailfish(text)
		snailfishs = append(snailfishs, snailfish)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return snailfishs
}

func readSnailfish(text string) ([]interface{}, int) {
	var snailfishs []interface{}
	length := 0

	for i := 1; i < len(text)-1; i++ {
		c := text[i]

		if length > 0 {
			continue
		} else if c == ',' {
			continue
		} else if c == '[' {
			snailfish, size := readSnailfish(text[i:])
			snailfishs = append(snailfishs, snailfish)
			i += size - 1
		} else if c == ']' {
			length = i
		} else {
			number, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}

			snailfishs = append(snailfishs, number)
		}
	}

	return snailfishs, length + 1
}

func magnitude(snailfish []interface{}) int {
	magnitudeLeft := 0
	magnitudeRight := 0

	if ok, number := isANumber(snailfish[0]); ok {
		magnitudeLeft = number
	} else if ok, array := isAnArray(snailfish[0]); ok {
		magnitudeLeft = magnitude(array)
	} else {
		panic("Noooooooo left")
	}

	if ok, number := isANumber(snailfish[1]); ok {
		magnitudeRight = number
	} else if ok, array := isAnArray(snailfish[1]); ok {
		magnitudeRight = magnitude(array)
	} else {
		panic("Noooooooo right")
	}

	return 3*magnitudeLeft + 2*magnitudeRight
}

// ### REDUCE

func reduce(snailfish interface{}) interface{} {
	actionDone := true

	for actionDone {
		actionDone = false
		snailfish, actionDone, _, _ = explode(0, snailfish)

		if !actionDone {
			snailfish, actionDone = split(snailfish)

		}
	}

	return snailfish
}

// ### EXPLODE

func explode(deep int, snailfishs interface{}) (interface{}, bool, int, int) {
	if ok, number := isANumber(snailfishs); ok {
		return number, false, 0, 0
	}

	ok, array := isAnArray(snailfishs)

	if !ok {
		panic("Should never happens")
	}

	if deep >= 4 {
		if DEBUG {
			fmt.Printf("Explode the value %v\n", array)
		}

		explodedLeft := array[0].(int)
		explodedRight := array[1].(int)

		return 0, true, explodedLeft, explodedRight
	}

	explodedInLeft := false
	explodedInRight := false
	explodedLeft := 0
	explodedRight := 0
	array[0], explodedInLeft, explodedLeft, explodedRight = explode(deep+1, array[0])
	if explodedRight > 0 {
		array[1] = explodePropagationFromRight(array[1], explodedRight)
		explodedRight = 0
	}

	if !explodedInLeft {
		array[1], explodedInRight, explodedLeft, explodedRight = explode(deep+1, array[1])

	}

	if explodedLeft > 0 && !explodedInLeft {
		array[0] = explodePropagationFromLeft(array[0], explodedLeft)
		explodedLeft = 0
	}

	return array, explodedInRight || explodedInLeft, explodedLeft, explodedRight
}

func explodePropagationFromRight(snailfishs interface{}, value int) interface{} {
	if ok, number := isANumber(snailfishs); ok {
		return number + value
	} else if ok, array := isAnArray(snailfishs); ok {
		array[0] = explodePropagationFromRight(array[0], value)

		return array
	} else {
		panic("Should never happens")
	}
}

func explodePropagationFromLeft(snailfishs interface{}, value int) interface{} {
	if ok, number := isANumber(snailfishs); ok {
		return number + value
	} else if ok, array := isAnArray(snailfishs); ok {
		array[1] = explodePropagationFromLeft(array[1], value)

		return array
	} else {
		panic("Should never happens")
	}
}

// ### SPLIT

func split(snailfishs interface{}) (interface{}, bool) {
	if ok, array := isAnArray(snailfishs); ok {
		splitted := false
		array[0], splitted = split(array[0])

		if !splitted {
			array[1], splitted = split(array[1])
		}

		return array, splitted
	}

	ok, numberInt := isANumber(snailfishs)
	if !ok {
		panic("Should never happens")
	}

	number := float64(numberInt)

	if number >= 10 {
		array := []interface{}{int(math.Floor(number / 2)), int(math.Ceil(number / 2))}
		return array, true
	}

	return numberInt, false
}

// ### UTIL

func isANumber(snailfish interface{}) (bool, int) {
	if reflect.TypeOf(snailfish).Kind() == reflect.Int {
		return true, snailfish.(int)
	}

	return false, 0
}

func isAnArray(snailfish interface{}) (bool, []interface{}) {
	if reflect.TypeOf(snailfish).Kind() == reflect.Slice {
		return true, snailfish.([]interface{})
	}

	return false, []interface{}{}
}
