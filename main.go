package main

import (
	"github.com/andreposman/go-pokedex/internal/repl"
	terminal "github.com/andreposman/go-pokedex/internal/terminal"
)

func main() {
	terminal.PrintGreeting()
	repl.Start()
}
