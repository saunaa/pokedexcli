package main

import (
    "io/ioutil"
)

func PrintPokemonAscii(ascii_file string) string{
    b, err := ioutil.ReadFile(ascii_file)
    if err != nil {
        panic(err)
    }
    return (string(b))
}