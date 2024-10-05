package main

import (
	"context"
	"errors"

	"github.com/SplinterSword/GOGator/internal/database"
	"github.com/google/uuid"
)

func handleFollowing(s *State, cmd Command, user database.User) error {

	nullUUID := uuid.NullUUID{
		UUID:  user.ID,
		Valid: true,
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), nullUUID)
	if err != nil {
		WriteInTerminal(err.Error())
		return errors.New("failed to get following")
	}

	for _, feed := range following {
		WriteInTerminal(feed.FeedName)
	}
	return nil
}
