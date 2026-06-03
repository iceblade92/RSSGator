package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iceblade92/RSSGator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := state{cfg: &cfg}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("wrong amount of arguments")
		os.Exit(1)
	}

	commandName := os.Args[1]
	cmd := command{
		Name: commandName,
		Args: os.Args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
