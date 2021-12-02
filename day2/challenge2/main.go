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
	aim := 0

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

		fmt.Printf("(Horizontal : %d, Vertical : %d, Aim : %d) I want to %s with %d\n", horizontal, vertical, aim, direction, value)

		switch direction {
		case "forward":
			horizontal += value
			vertical += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Horizontal : %d. Vertical : %d\n", horizontal, vertical)
	fmt.Println(vertical * horizontal)
}
