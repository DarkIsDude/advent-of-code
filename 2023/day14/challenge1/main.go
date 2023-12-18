package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	platform := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		platform = append(platform, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	displatPlatform(platform)
	moveToTheTop(platform)
	fmt.Println("--------------------")
	displatPlatform(platform)

	fmt.Println(calculateTotalLoad(platform))
}

func calculateTotalLoad(platform []string) int {
	totalLoad := 0

	for i, line := range platform {
		for _, char := range line {
			if char == 'O' {
				totalLoad += len(platform) - i
			}
		}
	}

	return totalLoad
}

func displatPlatform(platform []string) {
	for _, line := range platform {
		fmt.Println(line)
	}
}

func moveToTheTop(platform []string) {
	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform); j++ {

			if platform[i][j] == 'O' {
				for k := i - 1; k >= 0; k-- {
					newK := -1

					if platform[k][j] == 'O' || platform[k][j] == '#' {
						newK = k + 1
					}

					if k == 0 && platform[k][j] == '.' {
						newK = k
					}

					if newK != -1 {
						fmt.Printf("Switch on column %d with line %d, %d\n", j, i, newK)
						platform[i] = platform[i][:j] + "." + platform[i][j+1:]
						platform[newK] = platform[newK][:j] + "O" + platform[newK][j+1:]
						break
					}
				}
			}
		}
	}
}
