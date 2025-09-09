package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.ResetUsers(context.Background()); err != nil {
		return err
	}

	if err := s.cfg.ClearUser(); err != nil {
		return err
	}

	fmt.Println("reset successful!")
	return nil
}
