package main

import(
	"fmt"
	"os"
)

func commandExit(s *state) error {
	fmt.Println("Leaving the Spiral... Goodbye!")
	os.Exit(0)
	return nil
}