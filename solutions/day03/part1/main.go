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
	for i := 0; i < len(backpacks) - 1; i++ {
		items := strings.Split(backpacks[i], "")
		
		firstCompartment := items[0:len(items)/2]
		secondCompartment := items[len(items)/2:(len(items))]
		
		for j := 0; j < len(firstCompartment); j++ {
			item := string(firstCompartment[j])
			
			matchFound := false
			for k := 0; k < len(secondCompartment); k++ {
				secondItem := string(secondCompartment[k])

				if(item == secondItem) {
					matchFound = true
					break
				}
			}

			if matchFound == true {
				sum += strings.Index(itemsAtlas, item) + 1
				break
			}
		}
	}

	fmt.Println(sum)
}
