package main

import (
		"strings"
		"bufio"
		"os"
		"fmt"
		"net/http"
		"encoding/json"
		"io"
		"github.com/saunaa/pokedexcli/internal/pokecache"
		"time"
		
	)

var cache *pokecache.Cache 

func main() {
	cache = pokecache.NewCache(5* time.Second)
	var argument string
	var commands map[string]cliCommand
	commands = map[string]cliCommand{
		"exit": {
			name:        	"exit",
			description: 	"Exit the Pokedex",
			callback:    	func(arg string) error {
							return commandExit(argument)
			},
			},
		
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback: 		func(arg string) error {
							return commandHelp(commands)
			},
			},
		"map": {
			name:			"map",
			description:	"Displays the next 20 location areas",
			callback:		func(arg string) error {
							return commandMap(argument)
			},
    		},
		"mapb": {
			name:			"mapb",
			description:	"Displays the previous 20 location areas",
			callback: 		func(arg string) error {
							return commandMapb(argument)
			},
			},
		"explore": {
			name:			"explore",
			description: 	"Lists Pokemon available at a given location",
			callback:		func(arg string) error {
							return commandExplore(argument)
			},
			},

	}
	pokeascii := PrintPokemonAscii()
	fmt.Println(pokeascii)
	scanner := bufio.NewScanner(os.Stdin)
	for {	
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			line := scanner.Text()
			cleanLine := cleanInput(line)
			if len(cleanLine) > 1 {
				argument = cleanLine[1]
			}
			if len(cleanLine) > 2 {
				fmt.Print("too many arguments")
				return
			}
			command := cleanLine[0]
			if _, ok := commands[command]; !ok {
				fmt.Println("Unknown command")
				continue
			}
			err := commands[command].callback(argument)
				if err != nil {
					fmt.Println(err)
				}

		}
		
	}
}

func cleanInput(text string) []string {
	lowerCase:= strings.ToLower(text)
	cleanText := strings.Fields(lowerCase)

	return cleanText

}

func commandExit(arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0) 
	return nil
}

func commandHelp(commands map[string]cliCommand) error { 
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n ")
	for _, command := range commands {
		output := fmt.Sprintf("%v: %v", command.name, command.description)
		fmt.Println(output)
	}	
	return nil

}


func makeRequest(direction string, db any) error{ 
		res, err := http.Get(client.URL)
		if err != nil {
			return err
		}
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("connection failed")
		
		cache.Add(direction, data)

		}
		err = json.Unmarshal(data, db) 
		if err!= nil {
			return err 
		}
		
		return nil 
	}

func printLocations(direction string) error{
	if val, ok := cache.Get(direction); ok {
		err := json.Unmarshal(val, &clients) 
		if err != nil {
			return err
		return nil
	}
}
	client.UpdateURL(direction)
	makeRequest(direction, &clients)
	for _, result := range clients.Results {
	fmt.Println(result.Name)
	}
	return nil
}

		
func (client *APIclient) UpdateURL(newURL string) {
	client.URL = newURL
}

func commandMap(arg string) error{
	if clients.Next == "" {
		clients.Next = client.URL
	}
	printLocations(clients.Next)
	return nil
}

func commandMapb(arg string) error{
	if clients.Previous == client.URL || clients.Previous == ""{
		return fmt.Errorf("you're on the first page")
	}
	
	printLocations(clients.Previous)
	return nil
}

func commandExplore(location string) error{
	locationUrl := client.URL + "/" + location
	if val, ok := cache.Get(locationUrl); ok {
		err := json.Unmarshal(val, &areas) 
		if err != nil {
			return err
		}
	} else {
		err := makeRequest(locationUrl, &areas) 
		if err != nil {
			return err
		}
	}
	for _, pokemon := range areas.Pokemon_encounters {
		fmt.Println(pokemon.Pokemon.Name)
		}
	
	return nil
}

	

type cliCommand struct {
	name		string
	description string
	callback	func(arg string) error
}

type config struct {
	Next		string
	Previous	string
	Results		[]struct {
					Name	string
					Url		string
	}
}

type APIclient struct {
	URL		string
}

type LocationArea struct {
	Pokemon_encounters []struct {
		Pokemon struct {
			Name 	string 
			URL  	string 
		}
	}
}


var client = &APIclient{
	URL: 		"https://pokeapi.co/api/v2/location-area",
	}
var clients = &config{}

var areas = &LocationArea{}