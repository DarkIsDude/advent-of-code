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
		springs := lineSplited[0]
	  indicationsAsString := lineSplited[1]

		indications := []int{}
		for _, indication := range strings.Split(indicationsAsString, ",") {
			i, err := strconv.Atoi(indication)

			if err != nil {
				log.Fatal(err)
			}

			indications = append(indications, i)
		}

		sumOfArrangements += numberOfArrangements(springs, indications)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumOfArrangements)
}

func numberOfArrangements(spring string, indications []int) int {
	fmt.Printf("---- %s, %v\n", spring, indications)

	numberOfArrangements := 0
	queue := []combinationType{{
		index: 0,
		indications: indications,
		final: strings.Clone(spring),
	}}

	for len(queue) > 0 {
		combination := queue[0]
		queue = queue[1:]

		remainingSpring := spring[combination.index:]

		if len(combination.indications) == 0 {
			if strings.IndexByte(remainingSpring, '#') == -1 {
				for i := combination.index; i < len(combination.final); i++ {
					combination.final = combination.final[:i] + "." + combination.final[i+1:]
				}

				fmt.Println(combination.final)

				numberOfArrangements++
			}

			continue
		}

		indication := combination.indications[0]

		if len(remainingSpring) == 0 {
			continue
		}

		if remainingSpring[0] != '#' {
			queue = append(queue, combinationType{
				index: combination.index + 1,
				indications: combination.indications,
				final: combination.final[:combination.index] + "." + combination.final[combination.index+1:],
			})
		}

		if len(remainingSpring) < indication {
			continue
		}

		if len(combination.indications) > 1 {
			if len(remainingSpring) < indication + 1 {
				continue
			}

			if remainingSpring[indication] == '#' {
				continue
			}
		}


		valueExtracted := remainingSpring[:indication]
		if strings.IndexByte(valueExtracted, '.') != -1 {
			continue 
		}

		nextIndex := combination.index + indication
		nextFinal := strings.Clone(combination.final)

		for i := combination.index; i < nextIndex; i++ {
			nextFinal = nextFinal[:i] + "#" + nextFinal[i+1:]
		}

		if len(combination.indications) > 1 {
			nextFinal = nextFinal[:nextIndex] + "." + nextFinal[nextIndex+1:]
			nextIndex++
		}

		queue = append(queue, combinationType{
			index: nextIndex,
			indications: combination.indications[1:],
			final: nextFinal,
		})
	}

	return numberOfArrangements
}
