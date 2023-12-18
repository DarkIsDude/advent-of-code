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


	for i := 0; i <= 1_000; i++ {
		moveToTheTop(platform)
		moveToTheWest(platform)
		moveToTheSouth(platform)
		moveToTheEast(platform)

		fmt.Println(i, calculateTotalLoad(platform))
	}

	/*
	loop := map[int]int{
		0:  65,
		1:  63,
		2:  68,
		3:  69,
		4:  69,
		5:  65,
		6:  64,
	}
	*/

	loop := map[int]int{
		0:94263,// No
		1:94278,
		2:94295,
		3:94312,// No
		4:94313,// No
		5:94315,// No
		6:94309,// No, too high
		7:94302,// No, too high
		8:94283,// No
		9:94269,
		10:94258,
		11:94253,
		12:94245,
		13:94255,
	}

	for i := 202; i <= 1_000_000_000; i++ {
		if i % 1_000_000 == 0 {
			fmt.Println(i, loop[(i+7)%14])
		}
	}
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

func moveToTheEast(platform []string) {
	for i := 0; i < len(platform); i++ {
		for j := len(platform) - 1; j >= 0; j-- {

			if platform[i][j] == 'O' {
				for k := j + 1; k < len(platform); k++ {
					newK := -1

					if platform[i][k] == 'O' || platform[i][k] == '#' {
						newK = k - 1
					}

					if k == len(platform)-1 && platform[i][k] == '.' {
						newK = k
					}

					if newK != -1 {
						platform[i] = platform[i][:j] + "." + platform[i][j+1:]
						platform[i] = platform[i][:newK] + "O" + platform[i][newK+1:]
						break
					}
				}
			}
		}
	}
}

func moveToTheWest(platform []string) {
	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform); j++ {

			if platform[i][j] == 'O' {
				for k := j - 1; k >= 0; k-- {
					newK := -1

					if platform[i][k] == 'O' || platform[i][k] == '#' {
						newK = k + 1
					}

					if k == 0 && platform[i][k] == '.' {
						newK = k
					}

					if newK != -1 {
						platform[i] = platform[i][:j] + "." + platform[i][j+1:]
						platform[i] = platform[i][:newK] + "O" + platform[i][newK+1:]
						break
					}
				}
			}
		}
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
						platform[i] = platform[i][:j] + "." + platform[i][j+1:]
						platform[newK] = platform[newK][:j] + "O" + platform[newK][j+1:]
						break
					}
				}
			}
		}
	}
}

func moveToTheSouth(platform []string) {
	for i := len(platform) - 1; i >= 0; i-- {
		for j := 0; j < len(platform); j++ {

			if platform[i][j] == 'O' {
				for k := i + 1; k < len(platform); k++ {
					newK := -1

					if platform[k][j] == 'O' || platform[k][j] == '#' {
						newK = k - 1
					}

					if k == len(platform)-1 && platform[k][j] == '.' {
						newK = k
					}

					if newK != -1 {
						platform[i] = platform[i][:j] + "." + platform[i][j+1:]
						platform[newK] = platform[newK][:j] + "O" + platform[newK][j+1:]
						break
					}
				}
			}
		}
	}
}
