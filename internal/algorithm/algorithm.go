package algorithm

import "stations/internal/parser"



func AlgoBFD(connections []parser.Connection, stations []parser.Station , start string, end string, numTrain int) ([]string, error) {

	graph := make(map[string][]string)
for _, conn := range connections {
	from := conn.From
	to := conn.To
	graph[from] = append(graph[from], to)
	graph[to] = append(graph[to], from)
}
	return graph[start], nil
}