package main

import (
	"context"
	"errors"

	"github.com/SplinterSword/GOGator/internal/database"
)

type authHandler func(s *State, cmd Command, user database.User) error

func middlerwareLoggedin(next authHandler) CommandHandler {
	return func(s *State, cmd Command) error {
		user, err := s.db.GetUser(context.Background(), s.CurrentConfig.CurrentUser)
		if err != nil {
			return errors.New("user doesn't exist")
		}
		return next(s, cmd, user)
	}
}
