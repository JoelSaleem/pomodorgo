package timers

import tea "github.com/charmbracelet/bubbletea"

type Timers struct{}

func NewTimers() *Timers {
	return &Timers{}
}

func (t Timers) View() string {
	return "Your timers"
}

func (t Timers) Update(msg tea.Msg) tea.Cmd {
	return nil
}
