package feeds

import (
	"context"
	"encoding/xml"
	"errors"
	"html"
	"io"
	"net/http"
)

func removeUnescapedStrings(feed *RSSFeed) {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := 0; i < len(feed.Channel.Item); i++ {
		item := &feed.Channel.Item[i]
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	request, err := http.NewRequest("GET", feedURL, nil)
	if err != nil {
		return nil, errors.New("failed to create request")
	}
	request.Header.Add("User-Agent", "gator")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.New("failed to fetch feed")
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("failed to read feed")
	}

	feed := RSSFeed{}
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, errors.New("failed to unmarshal feed")
	}

	removeUnescapedStrings(&feed)

	return &feed, nil
}
