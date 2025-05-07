package input

import (

	"fmt"
	"os"
	"strconv"
)


// Save arguments to variables
type InputArgs struct{
	NetworkPath  string
	StartStation string
	EndStation   string
	NumTrains	int
}


 
// Processes user's input and handles errors
func ProcessInput(args []string) (*InputArgs, error) {

	

	if len(os.Args) != 5 {
		return  nil, fmt.Errorf("incorrect number of command-line arguments: expected 4 (networkPath, startStation, endStation, numTrains), but got %d", len(os.Args)-1,)
	}

	// Save arguments to variables
	networkPath := args[1]
	startStation := args[2]
	endStation := args[3]
	numTrainsStr := args[4]


	// Validate user's input
	if networkPath == "" {
		return nil, fmt.Errorf("network path cannot be empty")
	}

	if startStation == "" {
		return nil, fmt.Errorf("start station cannot be empty")
	}

	if endStation == "" {
		return nil, fmt.Errorf("end station cannot be empty")
	}

	if numTrainsStr == "" {
		return nil, fmt.Errorf("number of trains cannot be empty")
	}

	// Convert numTrainStr to an integer
	var err error
	numTrains, err := strconv.Atoi(numTrainsStr)
	if err != nil {
		return nil, fmt.Errorf("error converting number of trains ('%s') to an integer: %w", numTrainsStr, err)
	}

	if numTrains < 1 {
		return nil, fmt.Errorf("number of trains must be a positive integer, got %d", numTrains)
	}

	if startStation == endStation {
		return nil, fmt.Errorf("start station ('%s') and end station ('%s') cannot be the same", startStation, endStation)
	}

	// All validations passed, create the InputArgs struct
	InputArgs := &InputArgs{
		NetworkPath:  networkPath,
		StartStation: startStation,
		EndStation:   endStation,
		NumTrains:    numTrains,
	}




	return InputArgs, nil

}





