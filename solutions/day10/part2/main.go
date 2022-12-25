package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	commands := strings.Split(string(input), "\n")

	width := 40
	height := 6

	// initialize rows
	var rows []string
	for i := 0; i < height; i++ {
		var row string
		for j := 0; j < width; j++ {
			row += "."
		}
		rows = append(rows, row)
	}
	fmt.Println(rows)

	position := 1
	for command := 0; command < len(commands); command++ {
		currentCommand := strings.Fields(commands[command])

		if currentCommand[0] == "addx" {
			register, err := strconv.Atoi(currentCommand[1])
			if err != nil {
				fmt.Println(err)
			}
			registerValues = append(registerValues, registerValues[len(registerValues)-1])
			registerValues = append(registerValues, register + registerValues[len(registerValues)-1])
		} else {
			registerValues = append(registerValues, registerValues[len(registerValues)-1])
		}
	}

	sum := registerValues[19] * 20 + registerValues[59] * 60 + registerValues[99] * 100 + registerValues[139] * 140 + registerValues[179] * 180 + registerValues[219] * 220

	fmt.Println(sum)
}