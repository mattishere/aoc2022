package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	input, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println(error)
	}

	pairs := strings.Split(string(input), "\n")

	sum := 0
	for i := 0; i < len(pairs); i++ {
		pair := pairs[i]

		elves := strings.Split(string(pair), ",")

		firstElfSectionsRange := strings.Split(string(elves[0]), "-")
		secondElfSectionsRange := strings.Split(string(elves[1]), "-")

		firstElfSections := sections(firstElfSectionsRange)
		secondElfSections := sections(secondElfSectionsRange)

		// For part 2:		part2(firstElfSections, secondElfSections)
		if part1(firstElfSections, secondElfSections) {
			sum++
		}
	}

	fmt.Println(sum)
}


func sections(array []string) []int {
	range1, err := strconv.Atoi(strings.Replace(array[0], "\r", "", -1))
	range2, err := strconv.Atoi(strings.Replace(array[1], "\r", "", -1))

	if err != nil {
		fmt.Println(err)
	}

	var sections []int
	for section := range1; section <= range2; section++ {
		sections = append(sections, section)
	}

	return sections
}

func part1(array1, array2 []int) bool {
	var commonElements []int

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				commonElements = append(commonElements, array2[j])
			}
		}
	}
	if reflect.DeepEqual(commonElements, array1) || reflect.DeepEqual(commonElements, array2) {
		return true
	}
	return false
}

func part2(array1, array2 []int) bool {
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				return true
			}
		}
	}
	return false
}