package main

import (
	"context"
	"fmt"
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
		fmt.Printf(" - %s\n", item.Title)
	}
}
