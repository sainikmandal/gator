package main

import (
	"context"
	"fmt"
	"github.com/sainikmandal/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: follow <feed-url>")
	}
	feedURL := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("could not find feed with url %s: %w", feedURL, err)
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not follow feed: %w", err)
	}

	fmt.Printf("User %s is now following feed %s\n", follow.UserName, follow.FeedName)
	return nil

}
