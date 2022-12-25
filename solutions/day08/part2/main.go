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
	rows := strings.Split(string(input), "\n")
	highestScenicScore := 0
	for row := 0; row < len(rows); row++ {
		if row != 0 || row != len(rows) - 1 {
			trees := strings.Split(strings.TrimSpace(rows[row]), "")
			for tree := 0; tree < len(trees); tree++ {
				if tree != 0 || tree != len(trees)-1 {
					currentTree, err := strconv.Atoi(trees[tree])
					if err != nil {
						fmt.Println(err)
					}
					leftDistance := LeftDistance(currentTree, trees[:tree])
					rightDistance := RightDistance(currentTree, trees[tree+1:])
					topDistance := TopDistance(currentTree, tree, row, rows[:row])
					bottomDistance := BottomDistance(currentTree, tree, row, rows[row+1:])

					currentScenicScore := leftDistance * rightDistance * topDistance * bottomDistance
					if currentScenicScore > highestScenicScore {
						highestScenicScore = currentScenicScore
					}
				}
			}
		}
	}
	fmt.Println(highestScenicScore)
}

func RightDistance(tree int, trees []string) int {
	distance := 0
	for i := range trees {
		currentTree, err := strconv.Atoi(trees[i])
		if err != nil {
			fmt.Println(err)
		}
		distance++
		if tree <= currentTree {
			break
		}
	}
	return distance
}
func LeftDistance(tree int, trees []string) int {
	distance := 0
	for i := len(trees)-1; i >= 0; i -= 1 {
		currentTree, err := strconv.Atoi(trees[i])
		if err != nil {
			fmt.Println(err)
		}
		distance++
		if tree <= currentTree {
			break
		}
	}
	return distance
}
func TopDistance(tree, treePos, rowPos int, rows []string) int {
	distance := 0
	for i := len(rows)-1; i >= 0; i-=1 {
		if i != rowPos {
			currentRow := strings.Split(rows[i], "")
			currentTree, err := strconv.Atoi(currentRow[treePos])
			if err != nil {
				fmt.Println(err)
			}
			distance++
			if tree <= currentTree {
				break
			}
		}
	}
	return distance
}
func BottomDistance(tree, treePos, rowPos int, rows []string) int {
	distance := 0
	for i := range rows {
		if i != rowPos {
			currentRow := strings.Split(rows[i], "")
			currentTree, err := strconv.Atoi(currentRow[treePos])
			if err != nil {
				fmt.Println(err)
			}
			distance++
			if tree <= currentTree {
				break
			}
		}
	}
	return distance
}