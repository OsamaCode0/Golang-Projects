package parser

import (
	"fmt"
	"stations/internal/input"
	"strings"
	
)


type Connection struct{

	From string
	To string

}



// Parses the network map
func ParseMap() ([]string, []Connection, error){

		// Get the scanned file from the input processing function
		scanner, file := input.ProcessInput()
		defer file.Close()

		var stations []string
		var connections []Connection
		inStations := false
		inConnections := false
	
		for scanner.Scan() {
			
	
			line := strings.TrimSpace(scanner.Text())
		
			if line == "" {
				continue
			}
			if strings.HasPrefix(line, "#") {
				continue
			}

			// Mark sections
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
					parts := strings.Split(line, "-")

					if len(parts) != 2 {
					return nil, nil, fmt.Errorf("invalid connection line: %q", line)
					}

					from := strings.TrimSpace(parts[0])
					to := strings.TrimSpace(parts[1])

					if from == "" || to == "" {
						return nil, nil, fmt.Errorf("invalid connection endpoints: %q", line)
					}

					connections = append(connections, Connection{From: from, To: to})
				}
					
				}
			}
	
		if err := scanner.Err(); err != nil {
			fmt.Println("Error while parsing:", err)
		}
	
		// Debug prints
		fmt.Println("Stations:", stations)
		fmt.Println("Connections:", connections)


	return stations, connections, nil
}

