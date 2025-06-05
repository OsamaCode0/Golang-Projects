package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"stations/internal/input"
)

type Station struct {
	Name   string
	Coords [2]int
}

type Connection struct {
	From string
	To   string
}

const maxStations = 10000

func ParseMap(inputArgs *input.InputArgs) ([]Station, []Connection, error) {
	networkMap, err := os.Open(inputArgs.NetworkPath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening network map file: %w", err)
	}
	defer networkMap.Close()

	scanner := bufio.NewScanner(networkMap)
	var stations []Station
	var connections []Connection

	inStations := false
	inConnections := false
	sawStations := false
	sawConnections := false
	seenStations := make(map[string]Station)
	coordMap := make(map[[2]int]string)
	startStation := inputArgs.StartStation
	endStation := inputArgs.EndStation

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		switch strings.ToLower(line) {
		case "stations:":
			if sawConnections {
				return nil, nil, fmt.Errorf("line %d: 'stations:' section after 'connections:'", lineNum)
			}
			inStations = true
			inConnections = false
			sawStations = true
			continue
		case "connections:":
			if !sawStations {
				return nil, nil, fmt.Errorf("line %d: 'connections:' before 'stations:'", lineNum)
			}
			inStations = false
			inConnections = true
			sawConnections = true
			continue
		}

		if inStations {
			parts := strings.Split(line, ",")
			if len(parts) != 3 {
				return nil, nil, fmt.Errorf("line %d: invalid station format", lineNum)
			}

			name := strings.TrimSpace(parts[0])
			xStr := strings.TrimSpace(parts[1])
			yStr := strings.TrimSpace(parts[2])

			if name == "" {
				return nil, nil, fmt.Errorf("line %d: station name cannot be empty", lineNum)
			}

			x, err := strconv.Atoi(xStr)
			if err != nil || x < 0 {
				return nil, nil, fmt.Errorf("line %d: invalid x coordinate", lineNum)
			}

			y, err := strconv.Atoi(yStr)
			if err != nil || y < 0 {
				return nil, nil, fmt.Errorf("line %d: invalid y coordinate", lineNum)
			}

			if _, exists := seenStations[name]; exists {
				return nil, nil, fmt.Errorf("line %d: duplicate station name '%s'", lineNum, name)
			}

			coords := [2]int{x, y}
			if existing, exists := coordMap[coords]; exists {
				return nil, nil, fmt.Errorf("line %d: duplicate coordinates for '%s' and '%s'", lineNum, name, existing)
			}

			if len(seenStations) >= maxStations {
				return nil, nil, fmt.Errorf("exceeded maximum stations limit of %d", maxStations)
			}

			station := Station{Name: name, Coords: coords}
			stations = append(stations, station)
			seenStations[name] = station
			coordMap[coords] = name

		} else if inConnections {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, nil, fmt.Errorf("line %d: invalid connection format", lineNum)
			}

			from := strings.TrimSpace(parts[0])
			to := strings.TrimSpace(parts[1])

			if from == "" || to == "" {
				return nil, nil, fmt.Errorf("line %d: empty station name in connection", lineNum)
			}

			if _, exists := seenStations[from]; !exists {
				return nil, nil, fmt.Errorf("line %d: station '%s' does not exist", lineNum, from)
			}

			if _, exists := seenStations[to]; !exists {
				return nil, nil, fmt.Errorf("line %d: station '%s' does not exist", lineNum, to)
			}

			// Normalize connection direction
			if from > to {
				from, to = to, from
			}

			// Check for duplicate connections
			for _, conn := range connections {
				normalizedFrom := conn.From
				normalizedTo := conn.To
				if normalizedFrom > normalizedTo {
					normalizedFrom, normalizedTo = normalizedTo, normalizedFrom
				}
				if from == normalizedFrom && to == normalizedTo {
					return nil, nil, fmt.Errorf("line %d: duplicate connection '%s-%s'", lineNum, from, to)
				}
			}

			connections = append(connections, Connection{From: from, To: to})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error scanning file: %w", err)
	}

	if !sawStations {
		return nil, nil, fmt.Errorf("missing 'stations:' section")
	}
	if !sawConnections {
		return nil, nil, fmt.Errorf("missing 'connections:' section")
	}

	if _, exists := seenStations[startStation]; !exists {
		return nil, nil, fmt.Errorf("start station '%s' not found", startStation)
	}
	if _, exists := seenStations[endStation]; !exists {
		return nil, nil, fmt.Errorf("end station '%s' not found", endStation)
	}
	if startStation == endStation {
		return nil, nil, fmt.Errorf("start and end station cannot be the same")
	}

	return stations, connections, nil
}