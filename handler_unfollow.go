package main

import (
	"context"
	"fmt"

	"github.com/CrymsonShadows/blog-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("unfollow a url")
	}

	url := cmd.args[0]

	err := s.db.DeleteFeedFollowWithUserAndURL(context.Background(), database.DeleteFeedFollowWithUserAndURLParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("error trying to delete feed follow record from database: %v", err)
	}

	return nil
}
