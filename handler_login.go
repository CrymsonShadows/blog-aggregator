package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login expects a username")
	}
	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err != nil {
		log.Fatalf("user: %v not registered, err: %v", cmd.args[0], err)
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to: %s\n", cmd.args[0])
	return nil
}
