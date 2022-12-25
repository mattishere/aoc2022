package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	type Move struct {
		opponentMove	string 
		win				int
		draw			int
		loss			int
	}

	// Scoring
	win := 6
	draw := 3

	rockScore := 1
	paperScore := 2
	scissorsScore := 3

	// Moves
	rock := Move{ "A", paperScore, rockScore, scissorsScore }
	paper := Move{ "B", scissorsScore, paperScore, rockScore }
	scissors := Move{ "C", rockScore, scissorsScore, paperScore }

	input, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println(error)
	}

	rounds := string(input)

	splitRounds := strings.Split(rounds, "\n")

	sum := 0;

	for i := 0; i < len(splitRounds); i++ {
		moves := strings.Split(splitRounds[i], " ")

		opponentMove := moves[0]
		outcome := moves[1]

		var opponentMoveType Move
		switch opponentMove {
			case "A": opponentMoveType = rock 
			case "B": opponentMoveType = paper
			case "C": opponentMoveType = scissors
			default: fmt.Println("Uhh something is not right here, check your input.txt")
		}

		switch outcome {
			case "X": sum += opponentMoveType.loss
			case "Y": sum += opponentMoveType.draw + draw
			case "Z": sum += opponentMoveType.win + win
			default: fmt.Println("Uhh something is not right here, check your input.txt")
		}
	}
	fmt.Println(sum)
} 
