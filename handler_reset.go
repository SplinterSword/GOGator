package main

import (
	"context"
	"errors"
)

func handleReset(s *State, cmd Command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return errors.New("reset failed")
	}

	WriteInTerminal("the database has been reset")
	return nil
}
