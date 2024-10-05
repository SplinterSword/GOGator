package main

import (
	"context"
	"errors"
	"time"

	"github.com/SplinterSword/GOGator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeeds(s *State, cmd Command, user database.User) error {
	args := cmd.Args

	if len(args) != 2 {
		return errors.New("invalid number of arguments")
	}

	feedName := args[0]
	feedURL := args[1]

	nullUUID := uuid.NullUUID{
		UUID:  user.ID,
		Valid: true,
	}
	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    nullUUID,
	}

	_, err := s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		WriteInTerminal(err.Error())
		return errors.New("feed already exists or failed to add feed")
	}

	feednullID := uuid.NullUUID{
		UUID:  feed.ID,
		Valid: true,
	}
	follow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    nullUUID,
		FeedID:    feednullID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), follow)
	if err != nil {
		WriteInTerminal(err.Error())
		return errors.New("failed to follow feed")
	}

	return nil
}
