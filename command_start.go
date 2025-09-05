package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

func commandStart() error {
	fmt.Println("")
	fmt.Println("Let's begin young wizard! I am thinking of a spell...")
	fmt.Println()
	
	reader := bufio.NewScanner(os.Stdin)
	for {
		reader.Scan()
		guess := strings.ToLower(reader.Text())
		fmt.Printf("You guessed %s", guess)
		fmt.Println()
		break
	}
	return nil

}