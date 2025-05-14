package parser

import (
	"bufio"
	"fmt"
	"os"
	"stations/internal/input"
	"strings"
	"strconv"
)

type Station struct {
	Name string
	Coords [2]int
}

type Connection struct {
	From string
	To   string
}

const maxStations = 10000

// Parses the network map file
// Takes *InputArgs as argument and returns stations and []Connection that will be used as input to the pathfinder
func ParseMap(inputArgs *input.InputArgs) ([]Station, []Connection, error) {


	// Open the network map file.
	networkMap, err := os.Open(inputArgs.NetworkPath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening network map file '%w", err)
	}
	
	// Create a new Scanner
	scanner := bufio.NewScanner(networkMap)
	// Close the file after return
	defer networkMap.Close()


	// Declare slices of Station and Connection
	var stations []Station
	var connections []Connection

	
	
	// initialize variables
	inStations := false
	inConnections := false
	sawStations := false
	sawConnections := false
	seenStations := make(map[string]Station)
	startStation := inputArgs.StartStation
	endStation := inputArgs.EndStation



	// Loop through the file
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		// skip comments and empty lines
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Mark sections
		switch strings.ToLower(line) {
		case "stations:":
			// Check if sections are in order
			if sawConnections {
				return nil, nil, fmt.Errorf(`"stations:" section found after "connections:" section`)
			}
			inStations = true
			inConnections = false
			sawStations = true
			continue
		case "connections:":
			// Check that sections are in order and `stations:` was seen
			if !sawStations {
				return nil, nil, fmt.Errorf(`"connections:" section found before "stations:" section`)
			}
			inStations = false
			inConnections = true
			sawConnections = true
			continue
		}

		// stations: section logic
		if inStations {
			if line != "" {
				// Split each station to 3 parts(name, x-Coords, y-Coords)
				parts := strings.Split(line, ",")
				if len(parts) != 3 {
					return nil, nil, fmt.Errorf("invalid station name or coordinates: %q", line)
				}


				// Station name will be at parts[0]
				// trim each part
				stationName := strings.TrimSpace(parts[0])
				xStr := strings.TrimSpace(parts[1])
				yStr := strings.TrimSpace(parts[2])

				// Convert coordinates to integers and handle errors
				xCoords,err := strconv.Atoi(xStr)
				if err != nil{
					return nil,nil, fmt.Errorf("error converting X coordinate '%s' for station %q: %w", parts[1], stationName, err)
				}
				yCoords,err := strconv.Atoi(yStr)
				if err != nil{
					return nil,nil, fmt.Errorf("error converting Y coordinate '%s' for station %q: %w", parts[2], stationName, err)
				}

				// Create the Station struct
				newStation := Station{
					Name:   stationName,
					Coords: [2]int{xCoords, yCoords},
				}

			
				// Check for duplicates
				if _, isDup := seenStations[stationName]; isDup {
					return nil, nil, fmt.Errorf("duplicate station name %q", stationName)
				}
				// Check that the limit(10,000) of stations is not exeeced
				if len(seenStations) > maxStations {
					return nil, nil, fmt.Errorf("too many stations: limit is %d", maxStations)
				}

				// Store the station in the map
				seenStations[stationName] = newStation
				
				// Populate stations
				stations = append(stations, newStation)

			}
		// connections: section logic
		} else if inConnections {

			if line != "" {
				// Split each connection to two parts
				parts := strings.Split(line, "-")

				// Check that each connection contains two parts
				if len(parts) != 2 {
					return nil, nil, fmt.Errorf("invalid connection line: %q", line)
				}

				// Remove whitespace from both parts
				from := strings.TrimSpace(parts[0])
				to := strings.TrimSpace(parts[1])

				// Check that parts are not empty, and that the stations exists
				if from == "" || to == "" {
					return nil, nil, fmt.Errorf("invalid connection endpoints empty station name in connections section: %q", line)
				}
				if  _, exists := seenStations[from]; !exists{
					return nil, nil, fmt.Errorf(`"from" station "%s" doesn't exist in connection: %q`, from, line)
				}
				if  _, exists := seenStations[to]; !exists{
					return nil, nil, fmt.Errorf(`"to" station "%s"  doesn't exist in connection: %q`, to, line)
				}
				
				// Populate connections
				connections = append(connections, Connection{From: from, To: to})
			}
		}
	}

	// Handle scanning errors
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error scanning network map file: %w", err)
	}

	// Check that start and end stations exist
	if _, exists := seenStations[startStation]; !exists{
		fmt.Println(seenStations)
		return nil, nil, fmt.Errorf("start station '%s' not found in the network map stations", startStation)
	}
	if _, exists := seenStations[endStation]; !exists{
		return nil, nil, fmt.Errorf("end station '%s' not found in the network map stations", endStation)
	}

	
	// Make sure that both stations: and connections: sections exist
	if !sawStations {
		return nil, nil, fmt.Errorf(`missing "stations:" section`)
	}
	if !sawConnections {
		return nil, nil, fmt.Errorf(`missing "connections:" section`)
	}


	// Debug prints
	fmt.Println("Stations:", stations)
	fmt.Println("Connections:", connections)
	fmt.Println("number of trains:", inputArgs.NumTrains)

	return stations, connections, nil

}
