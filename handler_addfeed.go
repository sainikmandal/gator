package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sainikmandal/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("not enough args")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	currentUser := s.cfg.CurrentUserName
	if currentUser == "" {
		return fmt.Errorf("no current user set, please register or login first")
	}

	user, err := s.db.GetUserByName(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("could not find current user in DB: %w", err)
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	fmt.Println("Feed created:", feed)
	return nil
}
