package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iceblade92/RSSGator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("wrong amout of arguments")
	}

	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Println("User logged in:", cmd.Args[0])
	return nil

}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("wrong amout of arguments")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}
	fmt.Println("User registered:", cmd.Args[0])
	fmt.Printf("User created: %+v\n", user)
	return nil

}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUser(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset users: %w", err)
	}
	fmt.Println("Users Reset:")
	return nil

}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch users: %w", err)
	}
	fmt.Println("Users List:")
	for i := 0; i < len(users); i++ {
		if users[i].Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", users[i].Name)
		} else {
			fmt.Printf("* %s\n", users[i].Name)
		}

	}
	return nil

}

func handlerFetchFeed(s *state, cmd command) error {
	ctx := context.Background()
	result, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("RSSFeed: %+v\n", *result)
	return nil
}
