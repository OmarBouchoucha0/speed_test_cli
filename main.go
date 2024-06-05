package main

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

const timeout = time.Second * 5

type (
	errMsg error
)

type model struct {
	timer     timer.Model
	help      help.Model
	quitting  bool
	textInput textinput.Model
	err       error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Start Typing"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	m := model{
		timer:     timer.NewWithInterval(timeout, time.Millisecond),
		help:      help.New(),
		textInput: ti,
		err:       nil,
	}
	return m
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
	case errMsg:
		m.err = msg
		return m, nil

	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		cmds = append(cmds, cmd)

	case timer.TimeoutMsg:
		m.quitting = true
		return m, tea.Quit
	}

	m.textInput, _ = m.textInput.Update(msg)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	s := m.timer.View()
	if m.timer.Timedout() {
		s = "All done!"
	}
	s += "\n"
	if !m.quitting {
		s = "Exiting in " + s
	}
	return s + "\n" + fmt.Sprintf(
		"type something?\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
