package input

import (
	"fmt"
	"os"
	"strconv"
)

type InputArgs struct {
	NetworkPath  string
	StartStation string
	EndStation   string
	NumTrains    int
}

func ProcessInput(args []string) (*InputArgs, error) {
	if len(os.Args) != 5 {
		return nil, fmt.Errorf("incorrect number of arguments: expected 4, got %d", len(os.Args)-1)
	}

	networkPath := os.Args[1]
	startStation := os.Args[2]
	endStation := os.Args[3]
	numTrainsStr := os.Args[4]

	if networkPath == "" {
		return nil, fmt.Errorf("network path cannot be empty")
	}
	if startStation == "" {
		return nil, fmt.Errorf("start station cannot be empty")
	}
	if endStation == "" {
		return nil, fmt.Errorf("end station cannot be empty")
	}
	if startStation == endStation {
		return nil, fmt.Errorf("start and end station cannot be the same")
	}

	numTrains, err := strconv.Atoi(numTrainsStr)
	if err != nil {
		return nil, fmt.Errorf("invalid number of trains: %w", err)
	}
	if numTrains < 1 {
		return nil, fmt.Errorf("number of trains must be positive")
	}

	return &InputArgs{
		NetworkPath:  networkPath,
		StartStation: startStation,
		EndStation:   endStation,
		NumTrains:    numTrains,
	}, nil
}