package main

import (
	"fmt"
)

type Command interface {
	Execute()
}

type PrintCommand struct {
	message string
}

func (c *PrintCommand) Execute() {
	fmt.Println(c.message)
}

func main() {
	command := &PrintCommand{message: "Hello, world!"}

	command.Execute()
}
