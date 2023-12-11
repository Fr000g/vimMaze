package maze

import "fmt"

type Maze struct {
	Map [][]int
	Win bool
	x   int
	y   int
}

func MapDisplay(maze Maze) {
	for _, valueRow := range maze.Map {
		for _, valuePoint := range valueRow {
			if valuePoint == wall {
				fmt.Print("░░")
			} else if valuePoint == path {
				fmt.Print("  ")
			} else if valuePoint == person {
				fmt.Print("🧑")
			} else if valuePoint == end {
				fmt.Print("💰")
			}
		}
		fmt.Println()
	}
}

// control the dot to move in the maze
func (m *Maze) Goleft() {
	if m.Map[m.x][m.y-1] == path {
		m.Map[m.x][m.y] = path
		m.y--
		m.Map[m.x][m.y] = person
	} else if m.Map[m.x][m.y-1] == end {
		m.Win = true
	}
}

func (m *Maze) Goright() {
	if m.Map[m.x][m.y+1] == path {
		m.Map[m.x][m.y] = path
		m.y++
		m.Map[m.x][m.y] = person
	} else if m.Map[m.x][m.y+1] == end {
		m.Win = true
	}
}

func (m *Maze) GoUp() {
	if m.Map[m.x-1][m.y] == path {
		m.Map[m.x][m.y] = path
		m.x--
		m.Map[m.x][m.y] = person
	} else if m.Map[m.x-1][m.y] == end {
		m.Win = true
	}
}

func (m *Maze) GoDown() {
	if m.Map[m.x+1][m.y] == path {
		m.Map[m.x][m.y] = path
		m.x++
		m.Map[m.x][m.y] = person
	} else if m.Map[m.x+1][m.y] == end {
		m.Win = true
	}
}
