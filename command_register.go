package main

func (c *Commands) Register(name string, handler CommandHandler) {
	c.CommandMap[name] = handler
}
