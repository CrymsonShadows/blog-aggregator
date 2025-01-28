package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/CrymsonShadows/blog-aggregator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("was unable to get next feed to fetch from database")
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

	fmt.Printf("Feed: %s", feed.Channel.Title)
	for i, item := range feed.Channel.Item {
		fmt.Printf("Item %d: %s\n", i, item.Title)
	}
}
