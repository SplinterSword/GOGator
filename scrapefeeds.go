package main

import (
	"context"
	"errors"
	"time"

	"github.com/SplinterSword/GOGator/feeds"
	"github.com/SplinterSword/GOGator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *State) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		WriteInTerminal("failed to get next feed to fetch")
		return errors.New("failed to get next feed to fetch")
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		WriteInTerminal("failed to mark feed fetched: " + err.Error())
		return errors.New("failed to mark feed fetched")
	}

	feed, err := feeds.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	WriteInTerminal("Title: " + feed.Channel.Title)
	WriteInTerminal(" ")

	for i := 0; i < len(feed.Channel.Item); i++ {
		item := &feed.Channel.Item[i]

		FeedNullUUID := uuid.NullUUID{
			UUID:  nextFeed.ID,
			Valid: true,
		}

		PubTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			WriteInTerminal(err.Error())
			return err
		}

		post := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: PubTime,
			FeedID:      FeedNullUUID,
		}

		_, err = s.db.CreatePost(context.Background(), post)
		if err != nil && err.Error() != `pq: duplicate key value violates unique constraint "posts_url_key"` {
			WriteInTerminal(err.Error())
			return err
		}

		if err.Error() == `pq: duplicate key value violates unique constraint "posts_url_key"` {
			continue
		}

		if item.Title == "" {
			continue
		}

		WriteInTerminal("Title: " + item.Title)
	}

	WriteInTerminal(" ")
	WriteInTerminal(" ")
	return nil
}
