package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/CrymsonShadows/blog-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int = 2
	var err error = nil
	if len(cmd.args) >= 1 {
		limit, err = strconv.Atoi(cmd.args[1])
		if err != nil {
			return fmt.Errorf("error getting number for browsing limit: %v", err)
		}
	}

	posts, err := s.db.GetPostsByUserID(context.Background(), database.GetPostsByUserIDParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		fmt.Printf("error getting posts to browse: %v", err)
	}

	for i, post := range posts {
		fmt.Printf("Post %d: %s\n\t%s\n", i, post.Title.String, post.Url)
	}
	return nil
}
