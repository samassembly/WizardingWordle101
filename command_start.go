package main

import(
	"fmt"
	"os"
	"bufio"
	//"strings"
	"context"
	//"github.com/samassembly/WizardingWordle101/internal/database"
)

func commandStart(s *state) error {
	fmt.Println("")
	fmt.Println("Let's begin young wizard! I am thinking of a spell...")
	fmt.Println()
	
	reader := bufio.NewScanner(os.Stdin)
	for {
		reader.Scan()
		// guess := strings.ToLower(reader.Text())
		guess := reader.Text()

		//check if guess exists in db
		spell, err := s.db.GetSpell(context.Background(), guess)
		if err != nil {
			fmt.Printf("That spell has not been taught in class!")
			fmt.Printf(guess)
			continue
		}

		fmt.Printf("You guessed %s", spell.Name)
		fmt.Println()
		break
	}
	return nil
}