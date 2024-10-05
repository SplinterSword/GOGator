package main

import (
	"context"
	"errors"
	"time"

	"github.com/SplinterSword/GOGator/internal/database"
	"github.com/google/uuid"
)

func handleFollow(s *State, cmd Command, user database.User) error {
	args := cmd.Args

	if len(args) != 1 {
		return errors.New("invalid number of arguments")
	}

	url := args[0]

	feed, err := s.db.GetFeedFromURL(context.Background(), url)
	if err != nil {
		return errors.New("feed doesn't exist")
	}

	UsernullUUID := uuid.NullUUID{
		UUID:  user.ID,
		Valid: true,
	}

	FeednullUUID := uuid.NullUUID{
		UUID:  feed.ID,
		Valid: true,
	}

	follow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    FeednullUUID,
		UserID:    UsernullUUID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), follow)
	if err != nil {
		WriteInTerminal(err.Error())
		return errors.New("failed to follow feed")
	}

	WriteInTerminal(user.Name)
	WriteInTerminal("followed " + feed.Name)
	return nil
}
