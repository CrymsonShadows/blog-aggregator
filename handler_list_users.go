package main

import (
	"context"
	"fmt"

	"github.com/CrymsonShadows/blog-aggregator/internal/config"
)

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving users: %v", err)
	}

	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("error reading from config: %v", err)
	}

	for _, user := range users {
		if user.Name == cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
