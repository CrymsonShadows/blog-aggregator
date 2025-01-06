package main

import (
	"fmt"

	"github.com/CrymsonShadows/blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandsMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	_, ok := c.commandsMap[name]
	if !ok {
		c.commandsMap[name] = f
	}
}

func (c *commands) run(s *state, cmd command) error {
	if _, ok := c.commandsMap[cmd.name]; !ok {
		return fmt.Errorf("the %s command does not exist", cmd.name)
	}
	err := c.commandsMap[cmd.name](s, cmd)
	return err
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login expects a single argument, the username")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to: %s\n", cmd.args[0])
	return nil
}
