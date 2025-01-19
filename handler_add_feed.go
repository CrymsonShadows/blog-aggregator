package main

import (
	"context"
	"fmt"
	"time"

	"github.com/CrymsonShadows/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func hanlderAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("addfeed command requires a name and url")
	}
	name := cmd.args[0]
	url := cmd.args[1]
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user from database: %v", err)
	}

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name,
			Url:       url,
			UserID:    currentUser.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("error creating new feed in database: %v", err)
	}

	fmt.Println(feed)
	return nil
}
