package repl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/andreposman/go-pokedex/internal/model"
)

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}

func commandMap() error {
	var countInt = 1
	locationsURL := "https://pokeapi.co/api/v2/location/"
	var LocationResponse model.LocationResponse

	r, err := http.Get(locationsURL)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer r.Body.Close()

	errJSON := json.NewDecoder(r.Body).Decode(&LocationResponse)
	if errJSON != nil {
		log.Fatalf(errJSON.Error())
	}

	fmt.Println("\n-----------------------------------------")
	fmt.Printf("\nThese are the first 20 locations: \n\n")

	for _, v := range LocationResponse.Results {
		fmtLocation := strings.Title(strings.Replace(v.Name, "-", " ", -1))
		fmt.Printf("%v -  %s \n", countInt, fmtLocation)
		countInt++
	}

	fmt.Println()

	return nil
}
