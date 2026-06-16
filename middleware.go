package main

import (
	"context"

	"github.com/iceblade92/RSSGator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {
		loggedinuser := s.cfg.CurrentUserName
		resp, err := s.db.GetUser(context.Background(), loggedinuser)
		if err != nil {
			return err
		}

		return handler(s, cmd, resp)
	}
}
