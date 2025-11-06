package main

import (
)

func handlerPlay(s *state, cmd command) error {
	//start repl loop
	startRepl(s)
	return nil
}