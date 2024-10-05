package main

import (
	"context"
	"errors"

	"github.com/SplinterSword/GOGator/internal/database"
	"github.com/google/uuid"
)

func deleteFeedFollow(s *State, cmd Command, user database.User) error {
	url := cmd.Args[0]

	feed, err := s.db.GetFeedFromURL(context.Background(), url)
	if err != nil {
		WriteInTerminal(err.Error())
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

	feedFollow := database.DeleteFeedFollowParams{
		UserID: UsernullUUID,
		FeedID: FeednullUUID,
		Url:    url,
	}

	err = s.db.DeleteFeedFollow(context.Background(), feedFollow)
	if err != nil {
		WriteInTerminal(err.Error())
		return errors.New("failed to delete feed follow")
	}
	return nil
}
