package main

import (
	"fmt"
	"os"
	"stations/internal/algorithm"
	"stations/internal/input"
	"stations/internal/parser"
	"strings"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "Error: incorrect number of arguments")
		os.Exit(1)
	}

	inputArgs, err := input.ProcessInput(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	stations, connections, err := parser.ParseMap(inputArgs)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Calculate max path length based on network size
	maxPathLength := len(stations) * 2
	if maxPathLength > 100 {
		maxPathLength = 100 // Cap for large networks
	}

	allPaths, err := algorithm.FindAllPathsUpToLength(connections, inputArgs.StartStation, inputArgs.EndStation, maxPathLength)
if err != nil {
    fmt.Fprintln(os.Stderr, "Error:", err)
    os.Exit(1)
}

assignedPaths := algorithm.AssignPathsToTrains(allPaths, inputArgs.NumTrains)
simulationResult := algorithm.SimulateTrainMovement(assignedPaths, inputArgs.NumTrains)

// Print simulation results
for _, turn := range simulationResult {
    fmt.Println(strings.Join(turn, " "))
}
}