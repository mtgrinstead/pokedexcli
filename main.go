package main

import (
	"fmt"
    "bufio"
    "os"
    "os/exec"
    "strings"
)

// Names CliName and calls function to print it
var cliName = "Pokedex" 
func printPrompt() {
    fmt.Print(cliName, ">")
}

//prints if an unknown function is typed
func printUnknown() {
    fmt.Println(text, ": command not found")
}

type cliCommand struct {
    name        string
    description string
    callback    func() error
}



func main() {
	fmt.Println("Pokedex >")
}
