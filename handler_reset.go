package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	//drop user table
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to delete users: %w", err)
	}
	//drop spell table
	err = s.db.DeleteSpells(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to delete spells: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}