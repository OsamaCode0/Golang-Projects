package parser

import (
	"fmt"
	"stations/internal/input"
	"strings"
)

type Connection struct {
	From string
	To   string
}

const maxStations = 10000

// Parses the network map
func ParseMap() ([]string, []Connection, error) {

	// Get the scanned file from the input processing function
	scanner, file := input.ProcessInput()
	defer file.Close()

	var stations []string
	var connections []Connection
	inStations := false
	inConnections := false
	sawStations := false
	SawConnections := false
	seenStations := make(map[string]struct{}, maxStations)

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

		if inStations {
			if line != "" {
				if _, dup := seenStations[line]; dup {
					return nil, nil, fmt.Errorf("duplicate station %q", line)
				}
				seenStations[line] = struct{}{}

				if len(seenStations) > maxStations {
					return nil, nil, fmt.Errorf("too many stations: limit is %d", maxStations)
				}
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

	if !sawStations {
		return nil, nil, fmt.Errorf(`missing "stations:" section`)
	}
	if !SawConnections {
		return nil, nil, fmt.Errorf(`missing "connections:" section`)
	}

	// Debug prints
	fmt.Println("Stations:", stations)
	fmt.Println("Connections:", connections)
	fmt.Println("number of trains:", input.NumTrainsInt)

	return stations, connections, nil

}
