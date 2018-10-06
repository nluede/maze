package maze

import "fmt"

// Maze contains the structure of a maze and its size
type Maze struct {
	structure     [][]bool
	width, height int
}

// New creates a new maze without any passway
func New(width, height int) Maze {

	if width < 0 {
		width = 0
	}

	if height < 0 {
		height = 0
	}

	newMaze := make([][]bool, height)

	for i := range newMaze {
		newMaze[i] = make([]bool, width)
	}

	return Maze{newMaze, width, height}
}

// Set modifies the maze for a given coordinate. Coordinates start at 0,0 and go up to width-1, height-1
func (maze *Maze) Set(x, y int, isPassway bool) {
	if x >= maze.width || y >= maze.height || x < 0 || y < 0 {
		return
	}
	maze.structure[y][x] = isPassway
}

// IsPassway checks if there is a passwaz at the given coordinates or not. If there is a passway, true gets returned.
// If there is a wall at the given coordinates, false gets returned.
func (maze *Maze) IsPassway(x, y int) bool {
	return maze.structure[y-1][x-1]
}

// Print prints out the maze to stdout
func (maze *Maze) Print() {
	for _, row := range maze.structure {
		printRow(row)
	}
}

func printRow(row []bool) {
	passway := ".."
	wall := "██"

	for _, e := range row {
		if e {
			fmt.Print(passway)
		} else {
			fmt.Print(wall)
		}
	}

	fmt.Print("\n")
}
