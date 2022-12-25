package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println(error)
	}

	calories := string(input)

	caloriesPerElf := strings.Split(calories, "\n\n")

	var calculatedCalories []int
	for i := 0; i < len(caloriesPerElf); i++ {
		caloriesOfElf := caloriesPerElf[i]

		splitCalories := strings.Split(caloriesOfElf, "\n")

		sum := 0
		for j := 0; j < len(splitCalories); j++ {
			if splitCalories[j] != "" {
				caloriesInt, error := strconv.Atoi(splitCalories[j])

				if error != nil {
					fmt.Println(error)
				}

				sum = sum + caloriesInt

			}
		}

		calculatedCalories = append(calculatedCalories, sum)
	}

	sort.Ints(calculatedCalories)

	// For part 2: fmt.Println(calculatedCalories[len(calculatedCalories)-1] + calculatedCalories[len(calculatedCalories)-2] + calculatedCalories[len(calculatedCalories)-3])
	fmt.Println(calculatedCalories[len(calculatedCalories)-1])
}
