package main

import(
	"fmt"
)

func commandHelp(s *state) error {
	printAsciiBanner()
	fmt.Println("")
	fmt.Println("Welcome to the Spiral!")
	fmt.Println()
	fmt.Println("Test your knowledge from Ravenwood Academy!")
	fmt.Println()
	fmt.Println("Start a new round using the start command:")
	fmt.Println("    WizWord101 > start")
	fmt.Println()
	fmt.Println("Exit the application with the exit command:")
	fmt.Println("    WizWord101 > exit")
	fmt.Println()
	fmt.Println("Good luck young wizard!")
	return nil
}