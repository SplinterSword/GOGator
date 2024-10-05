package main

import (
	"context"
	"errors"
	"time"

	"github.com/SplinterSword/GOGator/internal/database"
	"github.com/google/uuid"
)

func handleRegister(s *State, cmd Command) error {
	args := cmd.Args

	if len(args) != 1 {
		return errors.New("username is required")
	}

	username := args[0]
	id := uuid.New()

	user := database.CreateUserParams{
		ID:        id,
		Name:      username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.CurrentConfig.SetUser(username)
	if err != nil {
		WriteInTerminal(err.Error())
		return errors.New("failed to login user")
	}

	_, err = s.db.CreateUser(context.Background(), user)
	if err != nil {
		WriteInTerminal(err.Error())
		return errors.New("user alread exists")
	}

	return nil
}
