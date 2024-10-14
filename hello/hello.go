package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Luchi", "Stunna", "Jakes"}
	// request a greeting message
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// print the message if no error was returned
	fmt.Println(messages)
}
