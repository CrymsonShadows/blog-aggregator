package main

import (
	"context"
	"fmt"

	"github.com/CrymsonShadows/blog-aggregator/internal/database"
)

func midllewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error getting current user from database: %v", err)
		}

		return handler(s, cmd, currentUser)
	}
}
