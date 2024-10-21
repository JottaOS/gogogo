package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)
	stacks := map[int][]string{}

	// scan initial crates
	for sc.Scan() {

		line := sc.Text()
		if(line == " 1   2   3   4   5   6   7   8   9 "){
			sc.Scan()
			break
		}

		chars := strings.Split(line, "")

		for i, char := range chars {
			if char != " " && char != "[" && char != "]" {
				trueIndex := i / 4
				stacks[trueIndex] = append(stacks[trueIndex], char)
			}
		}
	}

	// reversar todos los slices
	for _, stack := range stacks {
		slices.Reverse(stack)
	}

	for sc.Scan() {
		
		var amount, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &amount, &from, &to)
	
		// obtener stacks
		fromStack := stacks[from - 1]
		toStack := stacks[to - 1]
	
		// ahora ya puedo mover varios elementos a la vez, en el mismo orden
		toStack = append(toStack, fromStack[len(fromStack) - amount:]...)

		// eliminar los elementos appended al destino
		fromStack = fromStack[:len(fromStack) - amount]

		// reasignar los slices modificados al map
		stacks[from - 1] = fromStack
		stacks[to - 1] = toStack
	}

	// última iteración para concatenar el head de cada stack
	solution := ""
	 for i := 0; i < len(stacks); i++ {
		currentStack := stacks[i]
		solution += currentStack[len(currentStack) - 1]	
	}	

	// solución final
	fmt.Printf("%v\n", solution)
}
