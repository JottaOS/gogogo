package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

		var x1, x2, y1, y2 int

		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &x1, &x2, &y1, &y2)

		// check if first range overlaps at least one section with second range or vice versa
		if x1 <= y2 && x2 >= y1 {
			overlaps++
		}
	}

	fmt.Println(overlaps)
}
