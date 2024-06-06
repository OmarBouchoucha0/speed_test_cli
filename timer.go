package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

const timeout = time.Second * 5

type timerModel struct {
	timer    timer.Model
	help     help.Model
	quitting bool
}

func (m timerModel) InitTimer() tea.Cmd {
	return nil
}

func (m timerModel) UpdateTimer(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	case timer.TimeoutMsg:
		m.quitting = true
		return m, tea.Quit

	}
	return m, nil
}

func (m timerModel) ViewTimer() string {
	s := m.timer.View()

	if m.timer.Timedout() {
		s = "All done!"
	}
	s += "\n"
	if !m.quitting {
		s = "Exiting in " + s
		s += m.helpView()
	}
	return s
}
