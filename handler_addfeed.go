package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sainikmandal/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		now := time.Now()
		feed, err = s.db.CreateFeed(context.Background(), database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
			Name:      feedName,
			Url:       feedURL,
			UserID:    user.ID,
		})
		if err != nil {
			return fmt.Errorf("could not create feed: %w", err)
		}
		fmt.Printf("Created feed: %s (%s)\n", feed.Name, feed.Url)
	} else {
		fmt.Printf("Feed already exists: %s (%s)\n", feed.Name, feed.Url)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed follow: %w", err)
	}

	fmt.Printf("%s is now following %s\n", user.Name, feed.Name)

	return nil
}
