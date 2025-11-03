package main

import (
	"fmt"
	"log"
	"github.com/samassembly/wizardingwordle101/internal/config"
)

func main() {
	cfg, err := Config.Read()
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("sam")
	if err != nil {
		log.Fatalf("Failed to set current user: %v", err)
	}
	
	startRepl()
}