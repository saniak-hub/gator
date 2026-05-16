package main

import (
	"errors"

	"github.com/saniak-hub/gator/internal/config"
	"github.com/saniak-hub/gator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	v, ok := c.cmds[cmd.name]
	if ok {
		if err := v(s, cmd); err != nil {
			return err
		}
		return nil
	}
	return errors.New("Command not found")
}
