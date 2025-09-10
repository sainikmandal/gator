package main

import (
	"context"
	"fmt"

	"github.com/sainikmandal/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: unfollow <feed-url>")
	}

	feedURL := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return err
	}

	err = s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't unfollow feed: %w", err)
	}

	fmt.Printf("You have unfollowed feed %s\n", feedURL)

	return nil
}
