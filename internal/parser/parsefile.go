package parser

import (

	"fmt"
	"os"
	"stations/internal/input"
	"strings"
	"bufio"
	
)

type Connection struct {
	From string
	To   string
}

const maxStations = 10000

// Parses the network map
func ParseMap(inputArgs *input.InputArgs) ([]string, []Connection, error) {


	// Open the network map file.
	networkMap, err := os.Open(inputArgs.NetworkPath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening network map file '%w", err)
	}
	
	// Create a new Scanner
	scanner := bufio.NewScanner(networkMap)

	defer networkMap.Close()


	
	var stations []string
	var connections []Connection



	inStations := false
	inConnections := false
	sawStations := false
	SawConnections := false
	seenStations := make(map[string]bool, maxStations)
	startStation := inputArgs.StartStation
	endStation := inputArgs.EndStation

	// Loop through the file
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Mark sections
		switch strings.ToLower(line) {
		case "stations:":
			inStations = true
			inConnections = false
			sawStations = true
			continue
		case "connections:":
			inStations = false
			inConnections = true
			SawConnections = true
			continue
		}

		// stations: section logic
		if inStations {
			//startStation := inputs.StartStation
			//seenStart := false
			
			if line != "" {
				if _, dup := seenStations[line]; dup {
					return nil, nil, fmt.Errorf("duplicate station %q", line)
				}
				seenStations[line] = true
				if len(seenStations) > maxStations {
					return nil, nil, fmt.Errorf("too many stations: limit is %d", maxStations)
				}
				stations = append(stations, line)

			}
		// connections: section logic
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
				if !seenStations[parts[0]] || !seenStations[parts[1]]{
					return nil, nil, fmt.Errorf("station(s) doesn't exist: %q", line)
				}
				 
				connections = append(connections, Connection{From: from, To: to})
			}

		}

	}


	// Handle errors
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error scanning network map file: %w", err)
	}

	if _, exists := seenStations[startStation]; !exists{
		fmt.Println(seenStations)
		return nil, nil, fmt.Errorf("start station '%s' not found in the network map stations", startStation)
	}
	
	if _, exists := seenStations[endStation]; !exists{
		return nil, nil, fmt.Errorf("end station '%s' not found in the network map stations", endStation)
	}

	

	if !sawStations {
		return nil, nil, fmt.Errorf(`missing "stations:" section`)
	}
	if !SawConnections {
		return nil, nil, fmt.Errorf(`missing "connections:" section`)
	}


	// Debug prints
	fmt.Println("Stations:", stations)
	fmt.Println("Connections:", connections)
	fmt.Println("number of trains:", inputArgs.NumTrains)

	return stations, connections, nil

}
