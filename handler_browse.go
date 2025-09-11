package main

import (
	"context"
	"fmt"
	"github.com/sainikmandal/gator/internal/database"
	"strconv"
	"time"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	limit := 2 // default
	if len(cmd.Args) > 0 {
		n, err := strconv.Atoi(cmd.Args[0])
		if err == nil {
			limit = n
		}
	}

	posts, err := s.db.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("[%s] %s\n%s\n\n",
			post.PublishedAt.Time.Format(time.RFC822),
			post.Title,
			post.Url,
		)
	}

	return nil
}
