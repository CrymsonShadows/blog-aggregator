package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user: %v", err)
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows for user: %s: %v", s.cfg.CurrentUserName, err)
	}

	fmt.Printf("Feed follows for user: %s\n", s.cfg.CurrentUserName)
	for _, feedFollow := range feedFollows {
		fmt.Println(feedFollow.FeedName)
	}
	return nil
}
