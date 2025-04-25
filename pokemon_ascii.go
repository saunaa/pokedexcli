package main

import (
    "io/ioutil"
)

func PrintPokemonAscii() string{
    b, err := ioutil.ReadFile("pokemon_ascii.txt")
    if err != nil {
        panic(err)
    }
    return (string(b))
}