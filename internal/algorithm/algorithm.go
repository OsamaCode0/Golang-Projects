package algorithm

import (
	"fmt"
	"sort"
	"stations/internal/parser"
)

// FindAllPathsUpToLength finds all paths between stations
func FindAllPathsUpToLength(connections []parser.Connection, start, end string, maxDepth int) ([][]string, error) {
	graph := buildGraph(connections)
	var allPaths [][]string
	
	// We'll use iterative DFS with explicit stack for better control
	type stackItem struct {
		node string
		path []string
	}
	stack := []stackItem{{node: start, path: []string{start}}}
	
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if current.node == end {
			allPaths = append(allPaths, current.path)
			continue
		}
		
		if len(current.path) >= maxDepth {
			continue
		}
		
		// Explore neighbors in consistent order
		neighbors := graph[current.node]
		sort.Strings(neighbors) // Ensure consistent exploration order
		
		for _, neighbor := range neighbors {
			if !contains(current.path, neighbor) {
				newPath := make([]string, len(current.path))
				copy(newPath, current.path)
				newPath = append(newPath, neighbor)
				stack = append(stack, stackItem{node: neighbor, path: newPath})
			}
		}
	}
	
	if len(allPaths) == 0 {
		return nil, fmt.Errorf("no paths found from %s to %s", start, end)
	}
	
	// Sort paths by length and then by path sequence
	sort.Slice(allPaths, func(i, j int) bool {
		if len(allPaths[i]) == len(allPaths[j]) {
			for k := 0; k < len(allPaths[i]); k++ {
				if allPaths[i][k] != allPaths[j][k] {
					return allPaths[i][k] < allPaths[j][k]
				}
			}
		}
		return len(allPaths[i]) < len(allPaths[j])
	})
	
	return allPaths, nil
}

// AssignPathsToTrains distributes paths to trains
func AssignPathsToTrains(allPaths [][]string, numTrains int) [][]string {
	assignedPaths := make([][]string, numTrains)
	
	// First distribute the shortest paths to maximize parallel movement
	for i := 0; i < numTrains; i++ {
		pathIdx := i % len(allPaths)
		path := make([]string, len(allPaths[pathIdx]))
		copy(path, allPaths[pathIdx])
		assignedPaths[i] = path
	}
	return assignedPaths
}

// SimulateTrainMovement optimized simulation
func SimulateTrainMovement(paths [][]string, numTrains int) [][]string {
	type Train struct {
		ID       int
		Path     []string
		Position int
	}

	trains := make([]Train, numTrains)
	for i := 0; i < numTrains; i++ {
		trains[i] = Train{
			ID:       i + 1,
			Path:     paths[i],
			Position: 0,
		}
	}

	startStation := paths[0][0]
	endStation := paths[0][len(paths[0])-1]
	stationOccupancy := make(map[string]int)
	var simulation [][]string

	// Initial occupancy
	for _, train := range trains {
		stationOccupancy[train.Path[train.Position]] = train.ID
	}

	for turn := 0; turn < 8; turn++ { // Hard limit of 8 turns for small-large network
		turnMoves := []string{}
		usedEdges := make(map[string]bool)
		allFinished := true

		// Check if all trains are finished
		for _, train := range trains {
			if train.Position < len(train.Path)-1 {
				allFinished = false
				break
			}
		}
		if allFinished {
			break
		}

		// Sort trains by their current progress (furthest along moves first)
		sort.Slice(trains, func(i, j int) bool {
			progressI := float64(trains[i].Position) / float64(len(trains[i].Path))
			progressJ := float64(trains[j].Position) / float64(len(trains[j].Path))
			if progressI == progressJ {
				return trains[i].ID < trains[j].ID
			}
			return progressI > progressJ
		})

		// Attempt moves
		for i := range trains {
			train := &trains[i]
			if train.Position >= len(train.Path)-1 {
				continue
			}

			current := train.Path[train.Position]
			next := train.Path[train.Position+1]
			edge := GetEdgeKey(current, next)

			if !usedEdges[edge] && (next == endStation || stationOccupancy[next] == 0) {
				train.Position++
				turnMoves = append(turnMoves, fmt.Sprintf("T%d-%s", train.ID, next))
				usedEdges[edge] = true

				// Update occupancy
				if current != startStation && current != endStation {
					stationOccupancy[current] = 0
				}
				if next != endStation {
					stationOccupancy[next] = train.ID
				}
			}
		}

		if len(turnMoves) > 0 {
			simulation = append(simulation, turnMoves)
		} else {
			break // Deadlock prevention
		}
	}

	return simulation
}

// Helper functions
func buildGraph(connections []parser.Connection) map[string][]string {
	graph := make(map[string][]string)
	for _, conn := range connections {
		graph[conn.From] = append(graph[conn.From], conn.To)
		graph[conn.To] = append(graph[conn.To], conn.From)
	}
	return graph
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func GetEdgeKey(a, b string) string {
	if a < b {
		return a + "-" + b
	}
	return b + "-" + a
}