package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func Start() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				log.Fatal(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display the help command message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the CLI",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the name of the 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
	}
}
