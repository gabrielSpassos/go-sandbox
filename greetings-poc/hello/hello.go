package main

import (
    "fmt"
	"log"

    "example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

	// Request a greeting message.
    message, err := greetings.Hello("Gabriel")
	printSafeMessage(message, err)

	// A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}
	// Request greeting messages for the names.
    messages, errors := greetings.Hellos(names)
	printSafeMessages(messages, errors)

	// Request a greeting message.
    message, err = greetings.Hello("")
    printSafeMessage(message, err)
	
}

func printSafeMessages(messages map[string]string, err error) {
	// If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}

func printSafeMessage(message string, err error) {
	// If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}