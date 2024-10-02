package main

import (
	"context"
	"errors"
)

func handleLogin(s *State, cmd Command) error {
	args := cmd.Args

	if len(args) != 1 {
		return errors.New("username is required")
	}

	username := args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return errors.New("user doesn't exist")
	}

	err = s.CurrentConfig.SetUser(username)
	if err != nil {
		return err
	}

	WriteInTerminal("the user has been set")

	return nil
}
