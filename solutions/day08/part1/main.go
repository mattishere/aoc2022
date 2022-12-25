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
	visibleTrees := -1
	for row := 0; row < len(rows); row++ {
		if row == 0 || row == len(rows) - 1 {
			visibleTrees += len(rows[row])
		} else {
			trees := strings.Split(strings.TrimSpace(rows[row]), "")
			for tree := 0; tree < len(trees); tree++ {
				if tree == 0 || tree == len(trees)-1 {
					visibleTrees++
				} else {
					currentTree, err := strconv.Atoi(trees[tree])
					if err != nil {
						fmt.Println(err)
					}
					leftCheck := VisibleHorizontally(currentTree, trees[:tree])
					rightCheck := VisibleHorizontally(currentTree, trees[tree+1:])
					topCheck := VisibleVertically(currentTree, tree, row, rows[:row])
					bottomCheck := VisibleVertically(currentTree, tree, row, rows[row+1:])
					if leftCheck || rightCheck || topCheck || bottomCheck {
						visibleTrees++
					}
				}
			}
		}
	}
	fmt.Println(visibleTrees)
}

func VisibleHorizontally(tree int, trees []string) bool {
	for i := range trees {
		currentTree, err := strconv.Atoi(trees[i])
		if err != nil {
			fmt.Println(err)
		}
		if tree <= currentTree {
			return false
		}
	}
	return true
}
func VisibleVertically(tree, treePos, rowPos int, rows []string) bool {
	for i := range rows {
		if i != rowPos {
			currentRow := strings.Split(rows[i], "")
			currentTree, err := strconv.Atoi(currentRow[treePos])
			if err != nil {
				fmt.Println(err)
			}
			if tree <= currentTree {
				return false
			}
		}
	}
	return true
}