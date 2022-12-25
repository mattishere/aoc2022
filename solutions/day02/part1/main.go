package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	type Move struct {
		opponentMove	string 
		yourMove		string
		win				string 
		loss			string
		score			int
	}

	// Scoring
	win := 6
	draw := 3

	// Moves

	rock := Move{ "A", "X", "C", "B", 1 }
	paper := Move{ "B", "Y", "A", "C", 2 }
	scissors := Move{ "C", "Z", "B", "A", 3 }

	input, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println(error)
	}


	rounds := string(input)

	splitRounds := strings.Split(rounds, "\n")

	sum := 0;

	// We use len(splitRounds) - 1 because the last element is an empty array
	for i := 0; i < len(splitRounds) - 1; i++ {
		moves := strings.Split(splitRounds[i], " ")

		opponentMove := moves[0]
		yourMove := moves[1]

		var yourMoveType Move
		switch yourMove {
			case "X": yourMoveType = rock 
			case "Y": yourMoveType = paper
			case "Z": yourMoveType = scissors
			default: fmt.Println("Uhh something is not right here, check your input.txt")
		}

		switch opponentMove {
			case yourMoveType.opponentMove:
				sum += draw
			case yourMoveType.win:
				sum += win
			case yourMoveType.loss:
				sum += 0
			default: fmt.Println("Uhh something is not right here, check your input.txt")
		}

		sum += yourMoveType.score

	}

	fmt.Println(sum)
}
