package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/sainikmandal/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("username is required")
	}
	name := cmd.Args[0]

	_, err := s.db.GetUserByName(context.Background(), name)
	if err != nil {
		fmt.Println("user does not exists: ", name)
		os.Exit(1)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("Logged in as:", name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("username is required for registration")
	}
	name := cmd.Args[0]

	_, err := s.db.GetUserByName(context.Background(), name)
	if err == nil {
		fmt.Println("user with that name already exists")
		os.Exit(1)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("failed to set current user: %w", err)
	}

	fmt.Printf("User registered successfully:\nID: %s\nCreatedAt: %s\nUpdatedAt: %s\nName: %s\n",
		user.ID, user.CreatedAt, user.UpdatedAt, user.Name)

	return nil
}

func handlerListUsers(s *state, cmd command) error {
	currentUser := s.cfg.CurrentUserName

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return fmt.Errorf("No user found!")
	}

	for _, user := range users {
		if user == currentUser {
			fmt.Println(user + " (current)")
		} else {
			fmt.Println(user)
		}
	}

	return nil
}
