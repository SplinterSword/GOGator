package main

import (
	"context"
	"errors"
)

func handleReset(s *State, cmd Command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return errors.New("reset failed to delete users")
	}

	err = s.db.DeleteAllFeeds(context.Background())
	if err != nil {
		return errors.New("reset failed to delete feeds")
	}

	err = s.db.DeleteAllFeedFollows(context.Background())
	if err != nil {
		return errors.New("reset failed to delete feed follows")
	}

	err = s.db.DeleteAllPost(context.Background())
	if err != nil {
		return errors.New("reset failed to delete posts")
	}

	WriteInTerminal("the database has been reset")
	return nil
}
