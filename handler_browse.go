package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/SplinterSword/GOGator/internal/database"
	"github.com/google/uuid"
)

func handleBrowse(s *State, cmd Command, user database.User) error {
	var Limit int32 = 2
	WriteInTerminal("Enter the Number of feeds you want to browse (default is 2)")
	fmt.Scanf("%d", &Limit)

	UserNullUUID := uuid.NullUUID{
		UUID:  user.ID,
		Valid: true,
	}
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), UserNullUUID)
	if err != nil {
		WriteInTerminal(err.Error())
		return errors.New("failed to get feeds")
	}

	for _, feed := range feeds {
		getPost := database.GetPostForUserParams{
			FeedID: feed.FeedID,
			Limit:  Limit,
		}
		posts, err := s.db.GetPostForUser(context.Background(), getPost)
		if err != nil {
			WriteInTerminal(err.Error())
			return errors.New("failed to get posts")
		}

		for _, post := range posts {
			WriteInTerminal(post.Title)
			WriteInTerminal(post.Url)
			WriteInTerminal("")
		}

		WriteInTerminal(" ")
		WriteInTerminal(" ")
	}

	return nil
}
