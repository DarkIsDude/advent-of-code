package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	horizontal := 0
	vertical := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		action := scanner.Text()
		actionSplited := strings.Split(action, " ")
		direction := actionSplited[0]
		value, err := strconv.Atoi(actionSplited[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			horizontal += value
		case "down":
			vertical += value
		case "up":
			vertical -= value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(vertical * horizontal)
}
