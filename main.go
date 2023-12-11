package main

import "vimMaze/maze"

func main() {
	iMap := maze.GenerateMaze(20)
	maze.MapDisplay(iMap)
}
