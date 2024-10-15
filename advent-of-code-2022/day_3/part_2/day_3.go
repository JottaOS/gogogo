package main

// rucksacks...

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(file)

	prioritySum := 0

out:
	for {

		// read 3 lines
		chunk := []string{}
		for i := 0; i < 3; i++ {
			line, _, err := r.ReadLine()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break out
				}

				log.Fatal(err)
			}

			chunk = append(chunk, string(line))
		}

		// complexity: O(n^∞)
	badge:
		for _, firstElfItem := range chunk[0] {
			for _, secondElfItem := range chunk[1] {
				if firstElfItem == secondElfItem {
					for _, thirdElfItem := range chunk[2] {
						if firstElfItem == secondElfItem && secondElfItem == thirdElfItem {
							ascii := int(firstElfItem)
							// check if uppercase by ascii value (Z = 90)
							if ascii <= 90 {
								prioritySum += ascii - 38 // to match priority value
							} else {
								prioritySum += ascii - 96
							}

							break badge
						}
					}
				}
			}
		}
	}
	fmt.Printf("%v\n", prioritySum)
}
