package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/CrymsonShadows/blog-aggregator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("agg needs a time bewtween reqs")
	}

	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error parsing time duration: %v", err)
	}

	fmt.Printf("Collecting feeds every %v\n", cmd.args[0])
	ticker := time.NewTicker(time_between_reqs)
	for range ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("was unable to get next feed to fetch from database: %v", err)
	}

	currentTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:            nextFeed.ID,
		LastFetchedAt: currentTime,
	})
	if err != nil {
		fmt.Printf("error updating update time of feed in database: %v", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Printf("error fetching from %v", err)
	}

	fmt.Printf("Feed: %s\n", feed.Channel.Title)
	for i, item := range feed.Channel.Item {
		fmt.Printf("\tItem %d: %s\n", i, item.Title)
	}
}
