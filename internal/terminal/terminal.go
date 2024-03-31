package terminal

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func PrintGreeting() {

	name := "cli pokedex"
	version := "0.0.1"
	author := "André"

	greeting := figure.NewFigure(name, "", true)
	greeting.Print()

	fmt.Printf("\n  - Version: %s\n", version)
	fmt.Printf("  - Author: %s\n", author)
	fmt.Printf("\n------------------------------------------------------------------------------------\n\n")

}
