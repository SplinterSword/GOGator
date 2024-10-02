package main

import "context"

func handleGetUsers(s *State, cmd Command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.CurrentConfig.CurrentUser {
			WriteInTerminal("* " + user.Name + " (current)")
			continue
		}
		WriteInTerminal("* " + user.Name)
	}

	return nil
}
