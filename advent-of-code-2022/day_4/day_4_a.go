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

		elf1A, _ := strconv.Atoi(firstElfRange[0])
		elf1B, _ := strconv.Atoi(firstElfRange[1])
		elf2A, _ := strconv.Atoi(secondElfRange[0])
		elf2B, _ := strconv.Atoi(secondElfRange[1])

		// check if first range contains second range or vice versa 
		if elf1A >= elf2A && elf1B <= elf2B {
			overlaps++
		} else if elf2A >= elf1A && elf2B <= elf1B {
			overlaps++
		}

	}

	fmt.Println(overlaps)
}
