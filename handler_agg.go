package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Printf("Feed: %s\n", feed.Channel.Title)
	fmt.Printf("Description: %s\n\n", feed.Channel.Description)

	for _, item := range feed.Channel.Items {
		fmt.Printf("- %s\n", item.Title)
		fmt.Printf("  Link: %s\n", item.Link)
		fmt.Printf("  Description: %s\n\n", item.Description)
	}

	return nil
}
