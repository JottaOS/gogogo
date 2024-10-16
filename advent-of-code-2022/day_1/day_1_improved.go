package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	chunks := parse()

	result := []int{}
	for _, chunk := range chunks {
		result = append(result, sumChunk(chunk))
	}

	sort.Ints(result)

	fmt.Println(result[len(result)-1] + result[len(result)-2] + result[len(result)-3])
}

func sumChunk(chunk string) int {
	sum := 0
	elements := strings.Split(chunk, "\n")
	for _, element := range elements {
		number, _ := strconv.Atoi(element)
		sum += number
	}
	return sum
}

func parse() []string {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	chunks := []string{}
	chunk := ""
	for {
		line, _, err := r.ReadLine()

		if string(line) == "" {
			chunks = append(chunks, chunk)
			chunk = ""
		} else {
			chunk += fmt.Sprintf("%v\n", string(line))
		}

		if err != nil {
			break
		}
	}

	return chunks
}
