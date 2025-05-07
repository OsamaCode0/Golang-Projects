package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


// Save arguments to variables
var(
	networkPath = os.Args[1]
	startStation = os.Args[2]
	endStation = os.Args[3]
	numTrainsStr = os.Args[4] 
)

// Public Variables
var( 
	NumTrainsInt int 
)

 
// Processes user's input and handles errors
func ProcessInput() (*bufio.Scanner, *os.File, error){

	var err error

	if len(os.Args) != 5 {
		return nil, nil, fmt.Errorf(
			"incorrect number of command-line arguments: expected 4 (networkPath, startStation, endStation, numTrains), but got %d",len(os.Args)-1,
		 )
	}

	// Validate user's input
	if networkPath == "" {
		return nil, nil, fmt.Errorf("network path cannot be empty")
	}

	if startStation == "" {
		return nil, nil, fmt.Errorf("start station cannot be empty")
	}

	if endStation == "" {
		return nil, nil, fmt.Errorf("end station cannot be empty")
	}

	if numTrainsStr == "" {
		return nil, nil, fmt.Errorf("number of trains cannot be empty")
	}

	// Convert numTrainStr to an integer
	NumTrainsInt, err = strconv.Atoi(numTrainsStr)
	if err != nil {
		return nil, nil, fmt.Errorf("error converting number of trains ('%s') to an integer: %w", numTrainsStr, err)
	}

	if NumTrainsInt < 1 {
		return nil, nil, fmt.Errorf("number of trains must be a positive integer, got %d", NumTrainsInt)
	}

	if startStation == endStation {
		return nil, nil, fmt.Errorf("start station ('%s') and end station ('%s') cannot be the same", startStation, endStation)
	}



	// Open the network map file.
	networkMap, err := os.Open(networkPath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening network map file '%s': %w", networkPath, err)
	}

	// Create a new Scanner
	scanner := bufio.NewScanner(networkMap)
	

	return scanner, networkMap, nil

}





