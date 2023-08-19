package stats

import tea "github.com/charmbracelet/bubbletea"

type Stats struct{}

func NewStats() *Stats {
	return &Stats{}
}

func (s Stats) View() string {
	return "Your stats"
}

func (s Stats) Update(msg tea.Msg) tea.Cmd {
	return nil
}
