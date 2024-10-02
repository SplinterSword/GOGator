package main

type Command struct {
	Name string
	Args []string
}

type CommandHandler func(s *State, cmd Command) error

type Commands struct {
	CommandMap map[string]CommandHandler
}
