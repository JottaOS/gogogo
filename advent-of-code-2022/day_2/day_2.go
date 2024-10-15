package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
	Referencia:
	1	A X Rock
	2	B Y Paper
	3	C Z Scissors

	0   Loss
	3	Draw
	6	Win
*/

func main() {

	pointValues := map[string]int{"X": 1, "Y": 2, "Z": 3, "A": 1, "B": 2, "C": 3}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	r := bufio.NewReader(file)

	totalPoints := 0
	for {
		line, _, err := r.ReadLine()

		if err != nil {
			if(errors.Is(err, io.EOF)){
				fmt.Printf("%v\n", totalPoints)
				break
			}

			log.Fatal(err)
		}

		round := strings.Split(string(line), " ")

		fmt.Printf("%v\n", round)
		playerA := pointValues[round[0]]
		playerB := pointValues[round[1]]

		// check for a draw (3 points)
		if playerA == playerB {
			totalPoints += 3 + playerB
			continue
		}

		// if true, playerB (Me) wins (6 points)
		if (playerA+1)%3 == playerB%3 {
			totalPoints += playerB + 6
		} else {
			totalPoints += playerB
		}

	}

}
