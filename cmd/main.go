package main

import (
	"fmt"
	"os"
	"stations/internal/input"
	"stations/internal/parser"
	"stations/internal/algorithm"
)




func main() {
    inputArgs, err := input.ProcessInput(os.Args)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to validate user's input:", err)
        os.Exit(1)
    }

    stations, connections, err := parser.ParseMap(inputArgs)

		

    if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to parse network map:", err)
        os.Exit(1)
    }

		End, err := algorithm.AlgoBFD(connections, stations, inputArgs.StartStation, inputArgs.EndStation, inputArgs.NumTrains)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Doesnt work man", err)
			os.Exit(0)
		}

		for _, value := range End {
			fmt.Printf(" %s \n", value)
		}

}
