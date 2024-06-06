package main

// A simple program that opens the alternate screen buffer then counts down
// from 5 and then exits.

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"time"
)

type model struct{}

type tickMsg time.Time

func main() {
	p := tea.NewProgram(model(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "Hi. This program will is full screen..."
}
