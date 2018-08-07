package parseCommand

import "fmt"

type Command struct {
	Test string;
}

func ParseCommand() Command {
	fmt.Println("parse command placeholder 2")
	command := Command { Test: "hello" }
	return command
}