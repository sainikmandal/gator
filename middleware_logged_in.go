package main

import (
	"context"
	"fmt"
	"github.com/sainikmandal/gator/internal/database"
)

func middlewareLoggedIn(
	handler func(s *state, cmd command, user database.User) error,
) func(*state, command) error {
	return func(s *state, cmd command) error {
		if s.cfg.CurrentUserName == "" {
			return fmt.Errorf("no user logged in, please register or login first")
		}

		user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("could not find current user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
