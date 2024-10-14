package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	accumulator := 0
	var caloriesSlice []int

	for {
		line, _, err := reader.ReadLine()

		lineString := string(line)
		if lineString != "" {
			calories, err := strconv.Atoi(string(line))
			if err != nil {
				log.Fatal(err)
			}
			accumulator += calories
		} else {
			caloriesSlice = append(caloriesSlice, accumulator)
			accumulator = 0
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				log.Fatal(err)
			}
		}
	}

	// fmt.Printf("Final slice")
	// fmt.Printf("%v", totalCalories)

	maxCalories := slices.Max(caloriesSlice)
	fmt.Printf("%v\n", maxCalories)
}
