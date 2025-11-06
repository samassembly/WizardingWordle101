package main 

import(
	"context"
	"fmt"
	"time"
	"github.com/samassembly/WizardingWordle101/internal/database"
	"strconv"
)

func handlerSpell(s *state, cmd command) error {
	if len(cmd.Args) != 5 {
		return fmt.Errorf("Not enough spell attributes!")
	}

	name := cmd.Args[0]
	cost64, err := strconv.ParseInt(cmd.Args[1], 10, 32) //decimal, 32-bit
	if err != nil {
		return fmt.Errorf("Error converting cost to int32: %v", err)
	}
	cost := int32(cost64)
	school := cmd.Args[2]
	accuracy64, err := strconv.ParseInt(cmd.Args[3], 10, 32)
	if err != nil {
		return fmt.Errorf("Error converting accuracy to int32: %v", err)
	}
	accuracy := int32(accuracy64)
	spellType := cmd.Args[4]

	spell, err := s.db.CreateSpell(context.Background(), database.CreateSpellParams{
		Name:         name,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		Cost:         cost,
		SpellSchool:  school,
		Accuracy:     accuracy,
		SpellType:    spellType,
	})
	if err != nil {
		return fmt.Errorf("Failed to create Spell: %w", err)
	}

	fmt.Println("Spell saved successfully:")
	printSpell(spell)
	return nil
}

func printSpell(spell database.Spell) {
	fmt.Printf(" * Name:    %v\n", spell.Name)
	fmt.Printf(" * Cost:    %v\n", spell.Cost)
	fmt.Printf(" * SpellSchool:    %v\n", spell.SpellSchool)
	fmt.Printf(" * Accuracy:    %v\n", spell.Accuracy)
	fmt.Printf(" * SpellType:    %v\n", spell.SpellType)
}