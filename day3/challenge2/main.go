package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var binaries []string

	for scanner.Scan() {
		binaries = append(binaries, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	oxygen := oxygen(binaries)
	co2 := co2(binaries)

	fmt.Printf("Oxygen %d CO2 %d\n", oxygen, co2)

	fmt.Println(oxygen * co2)
}

func oxygen(binaries []string) int64 {
	binariesLength := len(binaries[0])
	for i := 0; len(binaries) > 1 && i < binariesLength; i++ {
		countDigit := 0
		for _, binary := range binaries {

			if binary[i] == '1' {
				countDigit += 1
			} else {
				countDigit -= 1
			}
		}

		var filteredBinaries []string
		for _, binary := range binaries {
			keep1 := countDigit >= 0 && binary[i] == '1'
			keep0 := countDigit < 0 && binary[i] == '0'
			if keep1 || keep0 {
				filteredBinaries = append(filteredBinaries, binary)
			}
		}

		binaries = filteredBinaries
	}

	if len(binaries) != 1 {
		panic("Should never happens")
	}

	oxygen, err := strconv.ParseInt(binaries[0], 2, 64)
	if err != nil {
		panic(err)
	}

	return oxygen
}

func co2(binaries []string) int64 {
	binariesLength := len(binaries[0])
	for i := 0; len(binaries) > 1 && i < binariesLength; i++ {
		countDigit := 0
		for _, binary := range binaries {

			if binary[i] == '1' {
				countDigit += 1
			} else {
				countDigit -= 1
			}
		}

		var filteredBinaries []string
		for _, binary := range binaries {
			keep1 := countDigit < 0 && binary[i] == '1'
			keep0 := countDigit >= 0 && binary[i] == '0'
			if keep1 || keep0 {
				filteredBinaries = append(filteredBinaries, binary)
			}
		}

		binaries = filteredBinaries
	}

	if len(binaries) != 1 {
		panic("Should never happens")
	}

	co2, err := strconv.ParseInt(binaries[0], 2, 64)
	if err != nil {
		panic(err)
	}

	return co2
}
