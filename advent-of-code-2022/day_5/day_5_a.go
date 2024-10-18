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

	sc := bufio.NewScanner(file)
	var stacks []string

	// primeras 8 líneas son los stacks iniciales
	for i := 0; i < 8; i++ {
		sc.Scan()
		line := sc.Text()
		fmt.Printf("%v\n", line)
		stacks = append(stacks, line)
	}

	// Saltar las siguientes dos lineas de relleno
	sc.Scan()
	sc.Scan()

	// leer movimientos
	var moves [][]int
	
	for sc.Scan() {
		var column, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &column, &from, &to)

		move := []int {column, from, to}
		moves = append(moves, move)
	}

	fmt.Printf("Moves: %v\n\n", moves)
	fmt.Printf("Stacks: %v\n", stacks)

}
