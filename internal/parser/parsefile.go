package parser

import (
	"fmt"
	"stations/internal/input"
	// strings
	
)


func ParseMap(){

	// Get the contents from the input processing function
	contentsByte := input.ProcessInput()


	// Search for stations: from the file's contents
	searchString := "stations:"
	searchBytes := []byte(searchString)
	n := len(searchBytes)
	found := false 

	for i := 0; i <= len(contentsByte)-n; i++ {
		if string(contentsByte[i:i+n]) == searchString {
			// Debug print
			fmt.Printf("Found '%s' at index %d\n", searchString, i)
			found = true
			break 
		}
		if found{
			
		}
	}

	if !found {
		fmt.Printf("%s not found in the network map content.\n", searchString)
	}




}

