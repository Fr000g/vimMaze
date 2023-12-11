package main

import (
	"os"
	"vimMaze/tui"
)

func main() {
	err := tui.StartTUI()
	if err != nil {
		os.Exit(1)
	}
}
