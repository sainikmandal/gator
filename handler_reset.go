package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	err := s.db.ResetUsers(context.Background())
	if err == nil {
		fmt.Println("reset successful!")
	}
	return err
}
