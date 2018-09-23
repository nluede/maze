package main

import "github.com/nluede/maze/imageprocession"

func main(){
	// TODO: errorhandling
	maze, _ := imageprocession.ReadImage("assets/2.png")
	maze.Print()
}
