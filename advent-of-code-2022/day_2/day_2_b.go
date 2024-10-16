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

	pointValues := map[string]int{"A": 1, "B": 2, "C": 3}
	results := map[string]string{"A": "C", "B": "A", "C": "B"}

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
			if errors.Is(err, io.EOF) {
				fmt.Printf("%v\n", totalPoints)
				break
			}

			log.Fatal(err)
		}

		round := strings.Split(string(line), " ")

		fmt.Printf("%v\n", round)
		playerA := round[0]
		result := round[1]

		// draw
		if result == "Y" {
			totalPoints += 3 + pointValues[playerA]
			continue
		}

		// defeat
		if result == "X" {
			playerB := results[playerA]
			totalPoints += pointValues[playerB]
			continue
		}

		// win
		if result == "Z" {
			playerB := results[results[playerA]]
			totalPoints += 6 + pointValues[playerB]
			continue
		}

	}

}
