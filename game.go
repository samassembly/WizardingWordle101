package main

import (
	"fmt"
	"context"
	"math/rand"
	//"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
)

//game table structure
type Game struct {
	NameAnswer		string
	CostAnswer		int32
	SchoolAnswer	string
	AccAnswer		int32
	TypeAnswer		string
	Table			[]string
}

//initialize game table
func newGame(s *state) (Game, error) {
	//get spells array
	spells, err := s.db.GetSpells(context.Background())
	if err != nil {
		// return Game, fmt.Errorf("Failed to read spellbook: %w", err)
	}
	//get random index
	spellsLength := len(spells)
	randInt := rand.Intn(spellsLength-1)
	//get spell at index
	spell := spells[randInt]
	//initialize game object
	newGame := Game{
		NameAnswer: 	spell.Name, 
		CostAnswer: 	spell.Cost, 
		SchoolAnswer: 	spell.SpellSchool, 
		AccAnswer: 		spell.Accuracy, 
		TypeAnswer: 	spell.SpellType, 
		// Table:			[]string{},	
	}
	headerStr := fmt.Sprintf("%-15s | %-6s | %-10s | %-8s | %-10s\n", "Name", "Cost", "School", "Accuracy", "Type")
	//lineStr := fmt.Sprintf(strings.Repeat("-", 15+6+10+8+10+12))
	newGame.Table = append(newGame.Table, headerStr)
	//newGame.Table = append(newGame.Table, lineStr)

	return newGame, nil
}

//update guess table

func addGuess(g *Game, n string, c int32, s string, a int32, t string) {
    var nString, cString, sString, aString, tString string

    if n == g.NameAnswer {
        nString = Green + "%-15s" + Reset
    } else {
        nString = Red + "%-15s" + Reset
    }

    if c == g.CostAnswer {
        cString = Green + "%-6d" + Reset
    } else {
        cString = Red + "%-6d" + Reset
    }

    if s == g.SchoolAnswer {
        sString = Green + "%-10s" + Reset
    } else {
        sString = Red + "%-10s" + Reset
    }

    if a == g.AccAnswer {
        aString = Green + "%-8d" + Reset
    } else {
        aString = Red + "%-8d" + Reset
    }

    if t == g.TypeAnswer {
        tString = Green + "%-10s" + Reset
    } else {
        tString = Red + "%-10s" + Reset
    }

    template := nString + " | " + cString + " | " + sString + " | " + aString + " | " + tString + "\n"
    guess := fmt.Sprintf(template, n, c, s, a, t)

    g.Table = append(g.Table, guess)
    //fmt.Printf("Added the guess:\n%v", guess)
}


//print guess table
func printGame(g Game) {
	for _, row := range g.Table {
		fmt.Printf("%v", row)
	}
	fmt.Println()
}