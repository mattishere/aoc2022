package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println(error)
	}

	itemsAtlas := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	backpacks := strings.Split(string(input), "\n")
	
	sum := 0
	for i := 0; i < len(backpacks) - 1; i+= 3 {
		var group []string


		group = backpacks[i:i+3]


		firstElf := strings.Split(group[0], "")
		secondElf := strings.Split(group[1], "")
		thirdElf := strings.Split(group[2], "")
		


		matchFound := false
		for j := 0; j < len(firstElf); j++ {
			item := string(firstElf[j])

			for k := 0; k < len(secondElf); k++ {
				secondItem := string(secondElf[k])

				for l := 0; l < len(thirdElf); l++ {
					thirdItem := string(thirdElf[l])

					if item == secondItem && item == thirdItem {
						matchFound = true 
						break
					}
				}

				if matchFound == true {
					sum += strings.Index(itemsAtlas, item) + 1
					break
				}
			}

			if matchFound == true {
				break
			}
		}
	
	}

	fmt.Println(sum)
}
