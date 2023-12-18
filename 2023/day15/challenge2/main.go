package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type lens struct {
	label string
	focal int
}

var boxes = make(map[int][]lens)

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		sequence := strings.Split(line, ",")

		for _, seq := range sequence {
			proceedSequence(seq)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	power := map[string]int{}

	for id, box := range boxes {
		for slot, lens := range box {
			power[lens.label] += power[lens.label] + (id + 1) * (slot + 1) * lens.focal
		}
	}

	sum := 0
	for label, power := range power {
		fmt.Printf("%s: %d\n", label, power)
		sum += power
	}

	fmt.Printf("Total power: %d\n", sum)
}

func proceedSequence(sequence string) {
	if sequence[len(sequence)-1] == '-' {
		processMinus(sequence)
	} else {
		proceedEqual(sequence)
	}
}

func processMinus(sequence string) {
	label := sequence[:len(sequence)-1]
	boxID := runHASHAlgotithm(label)

	indexToRemove := -1
	for i, lens := range boxes[boxID] {
		if lens.label == label {
			indexToRemove = i
		}
	}

	if indexToRemove != -1 {
		fmt.Printf("Removing %s from box %d\n", label, boxID)
		boxes[boxID] = append(boxes[boxID][:indexToRemove], boxes[boxID][indexToRemove+1:]...)
	}
}

func proceedEqual(sequence string) {
	sequenceSplitted := strings.Split(sequence, "=")
	label := sequenceSplitted[0]
	focalAsS := sequenceSplitted[1]
	focal, err := strconv.Atoi(focalAsS)

	if err != nil {
		log.Fatal(err)
	}

	boxID := runHASHAlgotithm(label)
	found := false

	for i, lens := range boxes[boxID] {
		if lens.label == label {
			fmt.Printf("Updating %s in box %d\n", label, boxID)
			boxes[boxID][i].focal = focal
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Adding %s to box %d\n", label, boxID)
		boxes[boxID] = append(boxes[boxID], lens{label: label, focal: focal})
	}
}

func runHASHAlgotithm(line string) int {
	hash := 0

	for _, char := range line {
		hash += int(char)
		hash *= 17
		hash = hash % 256
	}

	return hash
}
