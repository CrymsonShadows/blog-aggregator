package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CrymsonShadows/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("a name must be passed in to register")
	}
	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err == nil {
		log.Fatalf("user: %v already registered, err: %v", cmd.args[0], err)
	}

	user, err := s.db.CreateUser(context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.args[0],
		})
	if err != nil {
		return fmt.Errorf("error creating a user in the database: %v", err)
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User %s has been registered\n", cmd.args[0])
	fmt.Println(user)
	return nil
}
