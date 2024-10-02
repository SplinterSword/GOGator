package main

import (
	"context"
	"strconv"

	"github.com/SplinterSword/GOGator/feeds"
)

func handleAgg(s *State, cmd Command) error {
	url := cmd.Args[0]

	feed, err := feeds.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	WriteInTerminal("Title: " + feed.Channel.Title)
	WriteInTerminal("Link: " + feed.Channel.Link)
	WriteInTerminal("Description: " + feed.Channel.Description)

	for i := 0; i < len(feed.Channel.Item); i++ {
		item := &feed.Channel.Item[i]

		WriteInTerminal("Item " + strconv.Itoa(i+1))
		WriteInTerminal("Title: " + item.Title)
		WriteInTerminal("Link: " + item.Link)
		WriteInTerminal("Description: " + item.Description)
		WriteInTerminal("Published: " + item.PubDate)
	}

	return nil
}
