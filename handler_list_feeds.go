package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds from database: %v", err)
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting username for feed: %v", err)
		}

		fmt.Printf("Feed name: %s\nFeed url: %s\nUsername: %s\n", feed.Name, feed.Url, user.Name)
	}
	return nil
}
