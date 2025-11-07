package main

import(
	"fmt"
	"os"
	"bufio"
	"time"
	//"strings"
	"context"
	//"github.com/samassembly/WizardingWordle101/internal/database"
)

func commandStart(s *state) error {
	fmt.Println("")
	fmt.Println("Let's begin young wizard! I am thinking of a spell...")
	fmt.Printf(">")

	//initialize game
	game, err := newGame(s)
	if err != nil {
		return fmt.Errorf("Failed to start game: %w", err)
	} 
	
	reader := bufio.NewScanner(os.Stdin)
	for i := 0; i < 5; i++ {
		reader.Scan()
		// guess := strings.ToLower(reader.Text())
		guess := reader.Text()

		//check if guess exists in db
		spell, err := s.db.GetSpell(context.Background(), guess)
		if err != nil {
			fmt.Println("Try again! Remember the answer is case sensitive (It's leviOsa, not levioSA!)")
			fmt.Printf(">")
			continue
		}

		//add guess to table
		addGuess(&game, spell.Name, spell.Cost, spell.SpellSchool, spell.Accuracy, spell.SpellType)
		printGame(game)

		//check if guess is correct
		if spell.Name == game.NameAnswer && spell.Cost == game.CostAnswer && spell.SpellSchool == game.SchoolAnswer && spell.Accuracy == game.AccAnswer && spell.SpellType == game.TypeAnswer {
			fmt.Println("Well done young wizard! You are doing exceptionally well in your studies!")
			return nil
		}

		fmt.Printf(">")
	}
	fmt.Println("I'm sorry young wizard, those spells are not the spell I am thinking of.")
	time.Sleep(2 * time.Second)
	fmt.Printf("I was thinking of the spell %s\n", game.NameAnswer)
	time.Sleep(2 * time.Second)
	fmt.Println("Chin up, keep up your studies and you'll be a masterful wizard in no time!")
	return nil
}