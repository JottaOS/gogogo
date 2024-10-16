package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// return a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string) // make(map[key-type]value-type)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// Associate the retrieved message with the name
		messages[name] = message
	}

	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned message
// is selected at random
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi %v, Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
