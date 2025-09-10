package main

import (
	"context"
	"fmt"

	"github.com/sainikmandal/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not fetch feed follows: %w", err)
	}

	if len(follows) == 0 {
		fmt.Printf("User %s is not following any feeds.\n", user.Name)
		return nil
	}

	fmt.Printf("Feeds followed by %s:\n", user.Name)
	for _, f := range follows {
		fmt.Printf("- %s (%s)\n", f.FeedName, f.FeedID)
	}

	return nil
}
