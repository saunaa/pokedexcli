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
		"math/rand"
		"image"
		"bytes"
		"image/draw"
		
	)

var cache *pokecache.Cache



func main() {

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
		"catch":  {
			name: 			"catch",
			description:	"Throw a pokeball at a pokemon",
			callback:		func(arg string) error {
							return commandCatch(argument)
			},
			},
		"inspect": {
			name:			"inspect",
			description:	"Inspect a pokemon",
			callback:		func(arg string) error {
							return commandInspect(argument)
			},
			},
		"pokedex": {		
			name:			"pokedex",
			description:	"Shows the pokemon you caught",
			callback:		func(arg string) error {
							return commandPokedex(argument)
			},
		},

	}


	cache = pokecache.NewCache(30* time.Minute)

	pokeascii := PrintPokemonAscii("pokemon_ascii.txt")
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
				fmt.Println("too many arguments")
				continue
			}
			command := cleanLine[0]
			if _, ok := commands[command]; !ok {
				fmt.Println("Unknown command. Type help to list commands!")
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


func makeRequest(direction, request string, db any) error{ 
		res, err := http.Get(request)
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
	makeRequest(direction, client.URL, &clients)
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
		err := makeRequest(locationUrl,locationUrl, &areas) 
		if err != nil {
			return err
		}
	}
	for _, pokemon := range areas.Pokemon_encounters {
		fmt.Println(pokemon.Pokemon.Name)
		}
	
	return nil
}

func commandCatch(pokemonName string) error{
	pokemonUrl := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	var pokemon = &Pokemon{}
	throw := fmt.Sprintf("Throwing a Pokeball at %v...", pokemonName)
	fmt.Println(throw)
	pokeascii := PrintPokemonAscii("pokeball_ascii.txt")
	fmt.Println(pokeascii)
	if val, ok := cache.Get(pokemonUrl); ok{
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return err
		}
	} else {
		err := makeRequest(pokemonUrl, pokemonUrl, &pokemon)
		if err != nil {
			return err
		}
	}
	odds := int(pokemon.Base_experience/100)
	catch := rand.Intn(odds+1)
	if odds != catch {
		escaped := fmt.Sprintf("%v escaped!", pokemonName)
		fmt.Println(escaped)
		return nil
	}
	pokedex[pokemonName] = *pokemon
	cought := fmt.Sprintf("%v was cought!", pokemonName)
	fmt.Println(cought)
	return nil
}


func commandInspect(pokemonName string) error{
	pkm, ok := pokedex[pokemonName] 
	if !ok {
		return fmt.Errorf("you have not cought that pokemon")
	}
	if val, ok := cache.Get(pkm.Sprites.Front_default); ok {
		img, _, err := image.Decode(bytes.NewReader(val)) 
		if err != nil {
			return err
		}
		Convert_image(trimImage(img))
	}else {
		img, err := getImage(pkm.Sprites.Front_default)
		if err != nil {
			return err
		}
		Convert_image(trimImage(img))
	}
	name := fmt.Sprintf("Name: %v", pkm.Name)
	height := fmt.Sprintf("Height: %v", pkm.Height)
	weight := fmt.Sprintf("Weight: %v", pkm.Weight)
	fmt.Println("")
	fmt.Println(name)
	fmt.Println(height)
	fmt.Println(weight)
	fmt.Println("Stats:")
	for _, stat:= range pkm.Stats {
		pokemonstats := fmt.Sprintf("	-%v: %v", stat.Stat.Name, stat.Base_stat)
		fmt.Println(pokemonstats)
	}
	fmt.Println("Types:")
	for _, typ := range pkm.Types {
		pokemontypes := fmt.Sprintf("	- %v", typ.Type.Name)
		fmt.Println(pokemontypes) 

	}
	return nil
	
}

func commandPokedex(argument string) error{
	if len(pokedex) == 0 {
		return fmt.Errorf("You haven't caught any Pokemon")
	}
	fmt.Println("Your Pokedex:")
	for _, pkm := range pokedex {
		name := fmt.Sprintf("	-%v", pkm.Name)
		fmt.Println(name)	
	}
	return nil
}

func getImage(image_url string) (image.Image, error){
		res, err := http.Get(image_url)
		if err != nil {
			return nil, err
		}
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		res.Body.Close()
		if res.StatusCode > 299 {
			return nil, fmt.Errorf("connection failed")
		
		cache.Add(image_url, data)

		}
		img ,_ , err := image.Decode(bytes.NewReader(data)) 
		if err!= nil {
			return nil, err
		
		
		}
		return img, err

}

func trimImage(img image.Image) image.Image {
	bounds := img.Bounds()
	minX, minY, maxX, maxY := bounds.Max.X, bounds.Max.Y, bounds.Min.X, bounds.Min.Y

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a != 0 {
				if x < minX {
					minX = x
				}
				if y < minY {
					minY = y
				}
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}
	rect := image.Rect(minX, minY, maxX+1, maxY+1)
	croppedImg := image.NewRGBA(rect)
	draw.Draw(croppedImg, rect, img, image.Point{X: minX, Y: minY}, draw.Src)

	return croppedImg
}

var client = &APIclient{
	URL: 		"https://pokeapi.co/api/v2/location-area",
	}
var clients = &config{}

var areas = &LocationArea{}

var pokedex map[string]Pokemon

func init() {
	pokedex = make(map[string]Pokemon)
}