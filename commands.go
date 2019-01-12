package main

import (
	"github.com/bwmarrin/discordgo"

	"strings"
)

var (
	commMap = make(map[string]Command)

	help    = Command{"help", "Displays all commands, pretty obvious. Also can display specific information using `&help` and a command after, for example, `&help alembic`.", helpCommand}
	alembic = Command{"alembic", "Displays an alembic :)", alembicCommand}
)

// Command : Every command is made into a struct to make it simpler to work with and eliminate if statements
type Command struct {
	name        string
	description string
	exec        func(*discordgo.Session, *discordgo.MessageCreate)
}

func loadCommands() {
	commMap[help.name] = help
	commMap[alembic.name] = alembic
}

func parseCommand(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	if strings.Contains(string(command[0]), "&") { // If the first word of the message starts with `&`:`
		command = string(command[1:]) // Remove the `&` from the command
		if command == strings.ToLower(commMap[command].name) {
			commMap[command].exec(s, m)
		}
	}
	return
}
