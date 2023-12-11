package maze

import (
	"fmt"
	"math/rand"
)

const (
	wall = iota
	path
	person
	end
)

type Maze [][]int

func GenerateMaze(length int) (pureMap Maze) {
	odd := 0
	if length%2 == 0 {
		odd = length - 1
	} else {
		odd = length
	}

	var iMap [][]int
	for i := 0; i < odd; i++ {
		var row []int
		for j := 0; j < odd; j++ {
			if j%2 == 0 || i%2 == 0 {
				row = append(row, 0)
			} else {
				row = append(row, 1)
			}
		}
		iMap = append(iMap, row)
	}

	for i := 1; i < odd; i += 2 {
		for j := 1; j < odd; j += 2 {
			if i == odd-2 && j == odd-2 {
				break
			}
			if i == odd-2 {
				iMap[i][j+1] = 1
				continue
			}
			if j == odd-2 {
				iMap[i+1][j] = 1
				continue
			}
			if rand.Int()%2 == 0 {
				iMap[i][j+1] = 1
			} else {
				iMap[i+1][j] = 1
			}
		}
	}

	iMap[1][1] = person
	iMap[odd-2][odd-2] = end
	return iMap
}

func MapDisplay(maze [][]int) {
	for _, valueRow := range maze {
		for _, valuePoint := range valueRow {
			if valuePoint == wall {
				fmt.Print("##")
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
