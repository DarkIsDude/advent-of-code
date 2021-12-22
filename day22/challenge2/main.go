package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const FILE_LOCATION = "./input.txt"
const DEBUG = false

type Instruction struct {
	xStart int
	xEnd   int
	yStart int
	yEnd   int
	zStart int
	zEnd   int

	on bool
}

func main() {
	instructions := readFile()
	var instructionsApplied []Instruction

	for pos, instruction := range instructions {
		fmt.Printf("- Apply instruction %d\n", pos+1)
		instructionsApplied = applyInstruction(instruction, instructionsApplied)

		if DEBUG {
			fmt.Printf("- Instruction to do : %d\n", len(instructionsApplied))
		}
	}

	fmt.Println("-------------")

	cubes := 0
	for _, instruction := range instructionsApplied {
		if DEBUG {
			fmt.Printf("Apply instruction %d,%d %d,%d %d,%d %v\n", instruction.xStart, instruction.xEnd, instruction.yStart, instruction.yEnd, instruction.zStart, instruction.zEnd, instruction.on)
		}

		if instruction.on {
			cubes += volume(instruction)
		} else {
			cubes -= volume(instruction)
		}

		if DEBUG {
			fmt.Println(cubes)
		}
	}

	fmt.Println(cubes)
}

func readFile() []Instruction {
	var instructions []Instruction

	file, err := os.Open(FILE_LOCATION)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		if DEBUG {
			fmt.Printf("New instruction : %s\n", text)
		}

		r := regexp.MustCompile(`(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)
		matches := r.FindStringSubmatch(text)

		on := matches[1] == "on"

		startX, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		endX, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}

		startY, err := strconv.Atoi(matches[4])
		if err != nil {
			panic(err)
		}

		endY, err := strconv.Atoi(matches[5])
		if err != nil {
			panic(err)
		}

		startZ, err := strconv.Atoi(matches[6])
		if err != nil {
			panic(err)
		}

		endZ, err := strconv.Atoi(matches[7])
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, Instruction{
			xStart: startX,
			xEnd:   endX,
			yStart: startY,
			yEnd:   endY,
			zStart: startZ,
			zEnd:   endZ,

			on: on,
		})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return instructions
}

func applyInstruction(instruction Instruction, instructionsApplied []Instruction) []Instruction {
	var intersections []Instruction

	for _, applied := range instructionsApplied {
		if overlaps(instruction, applied) {
			if DEBUG {
				fmt.Printf("-- Overlaps detected\n")
			}
			intersection := getIntersection(applied, instruction)

			if applied.on == instruction.on {
				intersection.on = !applied.on
			} else {
				intersection.on = instruction.on
			}

			if DEBUG {
				fmt.Printf("-- Add intersection %d,%d %d,%d %d,%d %v\n", intersection.xStart, intersection.xEnd, intersection.yStart, intersection.yEnd, intersection.zStart, intersection.zEnd, intersection.on)
			}
			intersections = append(intersections, intersection)
		}
	}

	if instruction.on {
		instructionsApplied = append(instructionsApplied, instruction)
	}

	instructionsApplied = append(instructionsApplied, intersections...)

	return instructionsApplied
}

func overlaps(i1 Instruction, i2 Instruction) bool {
	if i1.xStart > i2.xEnd || i1.xEnd < i2.xStart {
		return false
	}

	if i1.yStart > i2.yEnd || i1.yEnd < i2.yStart {
		return false
	}

	if i1.zStart > i2.zEnd || i1.zEnd < i2.zStart {
		return false
	}

	return true
}

func getIntersection(old Instruction, new Instruction) Instruction {
	intersection := Instruction{}

	intersection.xStart = int(math.Max(float64(old.xStart), float64(new.xStart)))
	intersection.xEnd = int(math.Min(float64(old.xEnd), float64(new.xEnd)))

	intersection.yStart = int(math.Max(float64(old.yStart), float64(new.yStart)))
	intersection.yEnd = int(math.Min(float64(old.yEnd), float64(new.yEnd)))

	intersection.zStart = int(math.Max(float64(old.zStart), float64(new.zStart)))
	intersection.zEnd = int(math.Min(float64(old.zEnd), float64(new.zEnd)))

	return intersection
}

func volume(intersection Instruction) int {
	return int(math.Abs(float64(intersection.xStart-intersection.xEnd))+1) *
		int(math.Abs(float64(intersection.yStart-intersection.yEnd))+1) *
		int(math.Abs(float64(intersection.zStart-intersection.zEnd))+1)
}
