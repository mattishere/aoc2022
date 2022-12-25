package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	Before you say anything, this is really bad code, way too many for loops. Might go back and make it a bit prettier when I find the time!
*/

func main() {
	instructionsInput, err := os.ReadFile("instructions.txt")
	stacksInput, err := os.ReadFile("stacks.txt")

	if err != nil {
		fmt.Println(err)
	}

	rows := strings.Split(string(stacksInput), "\n")

	var rowsSplit [][]string
	for row := 0; row < len(rows) - 1; row++ {
		currRow := rows[row]

		var rowSplit []string
		for i := 1; i < len(currRow); i += 4 {
			rowSplit = append(rowSplit, string(currRow[i]))
		}

		rowsSplit = append(rowsSplit, rowSplit)
	}

	var stacks [][]string

	// Initialize stacks
	for i := 0; i < len(rowsSplit[0]); i++ {
		stacks = append(stacks, []string{})
	}

	// Add initial crates
	for row := 0; row < len(rowsSplit); row++ {
		currentRow := rowsSplit[row]

		for crate := 0; crate < len(currentRow); crate++ {
			if currentRow[crate] != " " {
				currentStack := stacks[crate]
				// Here we prepend instead of append, makes it a bit easier imo.
				stacks[crate] = append([]string{currentRow[crate]}, currentStack...)
			}
		}
	}

	instructions := strings.Split(string(instructionsInput), "\n")
	
	for instruction := 0; instruction < len(instructions); instruction++ {
		currentInstruction := strings.Fields(instructions[instruction])

		amount, err := strconv.Atoi(currentInstruction[1])
		from, err := strconv.Atoi(currentInstruction[3])
		to, err := strconv.Atoi(currentInstruction[5])

		if err != nil {
			fmt.Println(err)
		}

		result := move(stacks[from - 1], stacks[to - 1], amount)

		stacks[from - 1] = result[0]
		stacks[to - 1] = result[1]
	}

	var topCrates []string
	for stack := 0; stack < len(stacks); stack++ {
		currentStack := stacks[stack]
		topCrates = append(topCrates, currentStack[len(currentStack)-1])
	}

	fmt.Println(strings.Join(topCrates, ""))
}


func move(source, destination []string, amount int) [][]string {
	cargo := source[len(source)-amount:]
	source = source[:len(source)-amount]
	
	for crate := 0; crate < len(cargo); crate++ {
		// Part 1: destination = append(destination, cargo[len(cargo) - 1 - crate])
		destination = append(destination, cargo[crate])
	}

	return [][]string{source, destination}
}