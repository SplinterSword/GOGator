package main

import (
	"errors"
	"time"
)

func handleAgg(s *State, cmd Command) error {

	if len(cmd.Args) != 1 {
		WriteInTerminal("invalid number of arguments")
		return errors.New("invalid number of arguments")
	}

	time_between_req := cmd.Args[0]

	if time_between_req == "" {
		WriteInTerminal("time between requests is required")
		return errors.New("time between requests is required")
	}

	if time_between_req == "0" {
		WriteInTerminal("time between requests cannot be 0")
		return errors.New("time between requests cannot be 0")
	}

	if time_between_req[len(time_between_req)-1] != 's' && time_between_req[len(time_between_req)-1] != 'm' && time_between_req[len(time_between_req)-1] != 'h' {
		WriteInTerminal("time between requests must be in seconds, minutes or hours")
		return errors.New("time between requests must be in seconds, minutes or hours")
	}

	WriteInTerminal("Collecting Feeds in Every " + time_between_req)

	timeBetweenRequests, err := time.ParseDuration(time_between_req)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
