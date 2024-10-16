package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	var overlaps int

	for sc.Scan() {

		line := sc.Text()
		ranges := strings.Split(line, ",")

		firstElfRange := strings.Split(ranges[0], "-")
		secondElfRange := strings.Split(ranges[1], "-")

		x1, _ := strconv.Atoi(firstElfRange[0])
		x2, _ := strconv.Atoi(firstElfRange[1])
		y1, _ := strconv.Atoi(secondElfRange[0])
		y2, _ := strconv.Atoi(secondElfRange[1])

		// check if first range overlaps at least one section with second range or vice versa
		if (x1 >= y1 && x1 <= y2) || (x2 >= y1 && x2 <= y2) || (y1 >= x1 && y1 <= x2) || (y2 >= x1 && y2 <= x2) {
			overlaps++
		}

	}

	fmt.Println(overlaps)
}
