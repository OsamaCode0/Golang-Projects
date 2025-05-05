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
	NumTrainsStr = os.Args[4] 
)

// Public Variables
var( 
	NumTrainsInt int 
)

 
// Processes user's input and handles errors
func ProcessInput() (*bufio.Scanner, *os.File){

	var err error

	NumTrainsInt, err = strconv.Atoi(NumTrainsStr)
	if err != nil{
		fmt.Println("Error while converting trains to int: ", err)
		os.Exit(1)
	}


	// Check that correct amount of arguments are given
	if len(os.Args) != 5{
		fmt.Fprintln(os.Stderr, "Error: Too few command line arguments")
		os.Exit(1)
	}


	// Error handling
	if networkPath == ""{
		fmt.Fprintln(os.Stderr, "Error: Please enter path to the network map file")
		os.Exit(1)
	}
	if startStation == ""{
		fmt.Fprintln(os.Stderr, "Error: Please specify the starting station")
		os.Exit(1)
	}

	if endStation == ""{
		fmt.Fprintln(os.Stderr, "Error: Please specify the end station")
		os.Exit(1)
	}
	if NumTrainsStr == ""{
		fmt.Fprintln(os.Stderr, "Error: Please enter the amount of trains")
		os.Exit(1)
	}
	if NumTrainsInt < 1{
		fmt.Fprintln(os.Stderr, "Error: Number of trains has to be a positive integer")
	}

	if startStation == endStation{
		fmt.Fprintln(os.Stderr, "Error: Cannot have the same starting and ending station")
		os.Exit(1)
	}


	// Open networkmap for reading
	networkMap, err := os.Open(networkPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Network map not found!")
		os.Exit(1)
	}


	// Create a new Scanner
	scanner := bufio.NewScanner(networkMap)
	if err := scanner.Err(); err != nil {
		panic(err)
	}





	return scanner, networkMap

}





