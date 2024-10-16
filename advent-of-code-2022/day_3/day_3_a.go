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
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(file)

    prioritySum := 0
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal(err)
		}

		parsedLine := string(line)
		halfLength := len(parsedLine) / 2

		containers := [2]string{parsedLine[:halfLength], parsedLine[halfLength:]}

		// fmt.Printf("Original line: %v Length: %v\n %v  (%v)\t\t%v (%v) \n", parsedLine, len(parsedLine), containers[0], len(containers[0]), containers[1], len(containers[1]))

        // Lowercase item types a through z have priorities 1 through 26.
        // Uppercase item types A through Z have priorities 27 through 52.
        out:
        for _, leftChar := range containers[0] {
            for _, rightChar := range containers[1] {
                if(leftChar == rightChar){

                    ascii := int(leftChar)
                    // check if uppercase by ascii value (Z = 90)
                    if(ascii <= 90){
                        prioritySum += ascii - 38 // to match priority value
                    } else {
                        prioritySum += ascii - 96
                    }

                    break out
                }
            }

        }

	}
    fmt.Printf("%v\n", prioritySum)
}
