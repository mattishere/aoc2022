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

	var registerValues []int
	for command := 0; command < len(commands); command++ {
		currentCommand := strings.Fields(commands[command])

		if currentCommand[0] == "addx" {
			register, err := strconv.Atoi(currentCommand[1])
			if err != nil {
				fmt.Println(err)
			}
			if len(registerValues) == 0 {
				registerValues = append(registerValues, 1)
			}
			registerValues = append(registerValues, registerValues[len(registerValues)-1])
			registerValues = append(registerValues, register + registerValues[len(registerValues)-1])
		} else {
			if len(registerValues) == 0 {
				registerValues = append(registerValues, 1)
			}
			registerValues = append(registerValues, registerValues[len(registerValues)-1])
		}
	}

	sum := registerValues[19] * 20 + registerValues[59] * 60 + registerValues[99] * 100 + registerValues[139] * 140 + registerValues[179] * 180 + registerValues[219] * 220

	fmt.Println(sum)
}