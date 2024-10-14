package main

import (
	"fmt"
	"log"
	"rsc.io/quote"
)

func main() {
	var name string
	fmt.Print("Como te llamás?\n> ")
	var _, err = fmt.Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello " + name + "... Tengo un quote para ti...")
	fmt.Print(quote.Go())
}
