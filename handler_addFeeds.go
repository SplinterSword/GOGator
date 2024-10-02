package main

import (
	"context"
	"errors"

	"github.com/SplinterSword/GOGator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeeds(s *State, cmd Command) error {
	args := cmd.Args

	if len(args) != 2 {
		return errors.New("invalid number of arguments")
	}

	feedName := args[0]
	feedURL := args[1]
	username := s.CurrentConfig.CurrentUser

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return errors.New("user doesn't exist")
	}

	nullUUID := uuid.NullUUID{
		UUID:  user.ID,
		Valid: true,
	}
	feed := database.CreateFeedParams{
		Name:   feedName,
		Url:    feedURL,
		UserID: nullUUID,
	}

	_, err = s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		return errors.New("feed already exists or failed to add feed")
	}

	return nil
}
