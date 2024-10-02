package main

import "errors"

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.CommandMap[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}

	return handler(s, cmd)
}
