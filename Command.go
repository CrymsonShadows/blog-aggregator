package main

import (
	"fmt"

	"github.com/CrymsonShadows/blog-aggregator/internal/config"
	"github.com/CrymsonShadows/blog-aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	_, ok := c.registeredCommands[name]
	if !ok {
		c.registeredCommands[name] = f
	}
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.name]
	if !ok {
		return fmt.Errorf("the %s command does not exist", cmd.name)
	}
	return f(s, cmd)
}
