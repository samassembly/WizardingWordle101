package main

import(
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Leaving the Spiral... Goodbye!")
	os.Exit(0)
	return nil
}