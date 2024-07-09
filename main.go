package main

import (
    "log"
    "fmt"
    "github.com/mtgrinstead/pokedexcli/internal/pokeapi"
)

func main() {
    //startRepl()
    pokeapiClient := pokeapi.NewClient()

    resp, err := pokeapiClient.ListLocationAreas()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(resp)
}
