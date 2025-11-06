package main 

import(
	"context"
	"fmt"
	"time"
	"strings"
	"github.com/samassembly/WizardingWordle101/internal/database"
	"strconv"
)

func handlerSpell(s *state, cmd command) error {
	//input validation
	if len(cmd.Args) != 5 {
		return fmt.Errorf("Not enough spell attributes!")
	}

	//parse input, cast values
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

	//create spell in db
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

	//print confirmation
	fmt.Println("Spell saved successfully:")
	printSpell(spell)
	return nil
}

func handlerListSpells(s *state, cmd command) error {
	//get spells from db
	spells, err := s.db.GetSpells(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to retrieve spells: %w", err)
	}
	//print header
	fmt.Printf("%-15s | %-6s | %-10s | %-8s | %-10s\n", "Name", "Cost", "School", "Accuracy", "Type")
	fmt.Println(strings.Repeat("-", 15+6+10+8+10+12))
	//print spells
	for _, spell := range spells {
		fmt.Printf("%-15s | %-6d | %-10s | %-8d | %-10s\n", spell.Name, spell.Cost, spell.SpellSchool, spell.Accuracy, spell.SpellType)

	}
	return nil
}

func printSpell(spell database.Spell) {
	fmt.Printf(" * Name:    %v\n", spell.Name)
	fmt.Printf(" * Cost:    %v\n", spell.Cost)
	fmt.Printf(" * SpellSchool:    %v\n", spell.SpellSchool)
	fmt.Printf(" * Accuracy:    %v\n", spell.Accuracy)
	fmt.Printf(" * SpellType:    %v\n", spell.SpellType)
}