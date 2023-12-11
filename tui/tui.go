package tui

import (
	"strings"
	"vimMaze/maze"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	wall = iota
	path
	person
	end
	key
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
var brickStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#B6482F")).Render

type model struct {
	Maze maze.Maze
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+c":
			return m, tea.Quit
		case "j":
			m.Maze.GoDown()
			return m, nil
		case "k":
			m.Maze.GoUp()
			return m, nil
		case "h":
			m.Maze.Goleft()
			return m, nil
		case "l":
			m.Maze.Goright()
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.Maze = maze.GenerateMaze(min(msg.Height, msg.Width))
		return m, nil
	}

	var cmd tea.Cmd

	return m, cmd
}

func (m model) View() string {
	if m.Maze.Win == true {
		return "You win!🎉" + "\n\n  " + helpStyle("Press ctrl+c to exit")
	}
	var build strings.Builder
	for _, valueRow := range m.Maze.Map {
		for _, valuePoint := range valueRow {
			if valuePoint == wall {
				build.WriteString(brickStyle("░░"))
			} else if valuePoint == path {
				build.WriteString("  ")
			} else if valuePoint == person {
				build.WriteString("🧑")

			} else if valuePoint == end {
				build.WriteString("💰")
			}

		}

		build.WriteString("\n")
	}
	return build.String()
}

func StartTUI() error {
	p := tea.NewProgram(model{})
	_, err := p.Run()
	return err
}
