package main

import (
	//"fmt"
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
	"github.com/samassembly/WizardingWordle101/internal/config"
	"github.com/samassembly/WizardingWordle101/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

func main() {
	//load config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}

	//establish connection to db
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	
	//program state, used to pass connections to commands
	programState := &state{
		db: dbQueries,
		cfg: &cfg,
	}

	//registered commands in 'dev' mode, outside repl
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("play", handlerPlay)
	cmds.register("spell", handlerSpell)
	cmds.register("users", handlerListUsers)
	cmds.register("spells", handlerListSpells)

	//checking for argument existence
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	//run the provided command
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}