package main


import(

	"fmt"
	"os"
	"stations/internal/input"
	"stations/internal/parser"
	"stations/internal/pathfinder"

)




func main(){
	var err error

	_,_,err = input.ProcessInput()
	if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to validate user's input:", err)
        os.Exit(1)
    }

	_,_ ,err = parser.ParseMap()
	if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to parse network map:", err)
        os.Exit(1)
    }

	pathfinder.Dijkstra(input.NumTrainsInt)




}