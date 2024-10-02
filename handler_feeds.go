package main

import (
	"context"
	"errors"
)

func handleFeeds(s *State, cmd Command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return errors.New("failed to get feeds")
	}

	for _, feed := range feeds {
		WriteInTerminal(feed.Name)
		WriteInTerminal(feed.Url)
		user, err := s.db.GetUserFromID(context.Background(), feed.UserID.UUID)
		if err != nil {
			return errors.New("failed to get user in handleFeeds")
		}
		WriteInTerminal(user.Name)
	}
	return nil
}
