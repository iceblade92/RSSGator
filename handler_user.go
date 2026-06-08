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
