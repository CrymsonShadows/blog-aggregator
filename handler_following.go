package main

import (
	"context"
	"fmt"

	"github.com/CrymsonShadows/blog-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows for user: %s: %v", s.cfg.CurrentUserName, err)
	}

	fmt.Printf("Feed follows for user: %s\n", s.cfg.CurrentUserName)
	for _, feedFollow := range feedFollows {
		fmt.Println(feedFollow.FeedName)
	}
	return nil
}
