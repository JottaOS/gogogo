package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if(err != nil){
		log.Fatal(err)
	}

	// la fila tiene una sola línea
	sc := bufio.NewScanner(file)
	sc.Scan()
	characters := strings.Split(sc.Text(), "")	
	set := map[string]bool {}
	for i := range characters {
		// leer los siguientes 4 caracteres
		// espero que no haga falta validar los ultimos caracteres del archivo...
		chunk := characters[i:i+4]
		for _, char  := range chunk {
			set[char] = true
		}

		// chequear si tiene 4 caracteres distintos
		if(len(set) == 4) {
			fmt.Printf("%v\n", i + 4)
			break
		} else {
			// si no, reiniciar el set
			clear(set)
		}
	}
}