package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sainikmandal/gator/internal/database"
	"log"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <time_between_reqs>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	// Run immediately first, then on every tick
	for {
		scrapeFeeds(s)
		<-ticker.C
	}
}

func scrapeFeeds(s *state) {
	ctx := context.Background()

	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		fmt.Println("Error getting next feed:", err)
		return
	}

	err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		fmt.Println("Error marking feed as fetched:", err)
		return
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		fmt.Println("Error fetching feed:", err)
		return
	}

	fmt.Printf("Feed: %s\n", rssFeed.Channel.Title)

	for _, item := range rssFeed.Channel.Items {
		// Try to parse pubDate, fall back to now
		publishedAt, err := parseTime(item.PubDate)
		if err != nil {
			log.Printf("Error parsing time for %s: %v\n", item.Link, err)
			publishedAt = time.Now()
		}

		// Save post in DB
		err = s.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: sql.NullTime{Time: publishedAt, Valid: true},
			FeedID:      feed.ID,
		})
		if err != nil {
			// Ignore duplicate URLs, log other errors
			log.Printf("Error saving post (%s): %v\n", item.Link, err)
			continue
		}

		fmt.Printf(" - %s\n", item.Title)
	}
}

func parseTime(raw string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC3339,
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, raw); err == nil {
			return t, nil
		}
	}
	return time.Now(), fmt.Errorf("unknown time format: %s", raw)
}
