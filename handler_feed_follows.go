package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iceblade92/RSSGator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Function needs 1 argument")
	}

	user := s.cfg.CurrentUserName
	resp, err := s.db.GetUser(context.Background(), user)
	if err != nil {
		return err
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	createdFollow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    resp.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return err
	}
	fmt.Println(createdFollow.FeedName)
	fmt.Println(createdFollow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("No need for extra arguments")
	}

	user := s.cfg.CurrentUserName
	resp, err := s.db.GetUser(context.Background(), user)
	if err != nil {
		return err
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), resp.ID)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}

	return nil
}
