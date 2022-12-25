package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")

	length := 14

	if err != nil {
		fmt.Println(err)
	}

	characters := strings.Split(string(input), "")

	sum := 0
	for i := 0; i < len(characters) - 1 - length; i++ {
		currentStack := characters[i:i+length]

		var nonRepeated int
		for j := range(currentStack) {
			occurences := strings.Count(strings.Join(currentStack, ""), currentStack[j])
			if occurences == 1 {
				nonRepeated++
			} else {
				break
			}
		}

		sum++
		if nonRepeated == len(currentStack) {
			sum += len(currentStack) - 1
			break
		}
	}

	fmt.Println(sum)
}