package main


import(

	"fmt"
	"os"
	"stations/internal/input"
	"stations/internal/parser"

)




func main(){
	var err error

	input.ProcessInput()
	_,_ ,err = parser.ParseMap()
	if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to parse network map:", err)
        os.Exit(1)
    }



}