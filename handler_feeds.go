package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsWithUsers(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		fmt.Printf("Feed: %s\n", feed.FeedName)
		fmt.Printf("URL: %s\n", feed.FeedUrl)
		fmt.Printf("Added by: %s\n", feed.UserName)
		fmt.Println("-----------")
	}

	return nil
}
