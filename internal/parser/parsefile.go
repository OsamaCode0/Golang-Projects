package parser

import (

	"fmt"
	"stations/internal/input"
	
	
)



func ParseMap(){

		// Get the scanned file from the input processing function
		scanner, file := input.ProcessInput()
		defer file.Close()

	
		var stations []string
		var connections []string
		inStations := false
		inConnections := false
	
		for scanner.Scan() {
			line := scanner.Text()
	
			switch {
			case line == "stations:":
				inStations = true
				inConnections = false
				continue
			case line == "connections:":
				inStations = false
				inConnections = true
				continue
			}
	
			if inStations {
				if line != "" {
					stations = append(stations, line)
				}
			} else if inConnections {
				if line != "" {
					connections = append(connections, line)
				}
			}
		}
	
		if err := scanner.Err(); err != nil {
			fmt.Println("Error while parsing:", err)
		}
	
		// Debug prints
		fmt.Println("Stations:", stations)
		fmt.Println("Connections:", connections)


}

