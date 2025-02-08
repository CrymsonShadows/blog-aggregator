package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/CrymsonShadows/blog-aggregator/internal/database"
	"github.com/google/uuid"
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
		var title sql.NullString
		var description sql.NullString
		if item.Title == "" {
			title.Valid = false
		} else {
			title.Valid = true
			title.String = item.Title
		}
		if item.Description == "" {
			description.Valid = false
		} else {
			description.Valid = true
			description.String = item.Description
		}

		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			fmt.Printf("TIME: %v\n", item.PubDate)
			fmt.Printf("error parsing publish time for post: %s\n%v\n", item.Title, err)
			os.Exit(1)
		}

		fmt.Printf("\tItem %d: %s\n", i, item.Title)
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       title,
			Url:         item.Link,
			Description: description,
			PublishedAt: pubDate,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			if err.Error() == "pq: duplicate key value violates unique constraint \"posts_url_key\"" {
				continue
			}
			fmt.Printf("error creating post in database: %v\n", err)
			// os.Exit(1)
		}
	}
}
