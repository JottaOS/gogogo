package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func parseFile(file *os.File) ([]int, error) {

	reader := bufio.NewReader(file)

	accumulator := 0
	var caloriesSlice []int

	for {
		line, _, err := reader.ReadLine()

		lineString := string(line)
		if lineString != "" {
			calories, err := strconv.Atoi(string(line))
			if err != nil {
				return nil, err
			}
			accumulator += calories
		} else {
			caloriesSlice = append(caloriesSlice, accumulator)
			accumulator = 0
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				return caloriesSlice, nil
			} else {
				return nil, err
			}
		}
	}
}

func sum(numbers []int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func maxNElements(n int, slice []int) []int {
	maxValues := make([]int, n)
	for i := range maxValues {
		maxValues[i] = 0
	}

	for _, element := range slice {
		for i, topValue := range maxValues {
			if element > topValue {
				currentElement := element
				for i != len(maxValues) {
					aux := maxValues[i]
					maxValues[i] = currentElement
					currentElement = aux
					i++
				}
				break
			}
		}
	}

	return maxValues
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	caloriesPerElf, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	top3Elves := maxNElements(3, caloriesPerElf)

	fmt.Printf("%v\n", sum(top3Elves))
}
