package main

import (
	"fmt"
	"os"
	"stations/internal/input"
	"stations/internal/parser"

)




func main(){

	
	inputArgs, err := input.ProcessInput(os.Args)
	if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to validate user's input:", err)
        os.Exit(1)
    }
	_,_, err = parser.ParseMap(inputArgs)
	if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to parse network map:", err)
        os.Exit(1)
    }
	



}