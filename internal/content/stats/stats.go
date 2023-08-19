package stats

import (
	"github.com/JoelSaleem/pomodorgo/internal/content"
	tea "github.com/charmbracelet/bubbletea"
)

type Stats struct{}

func NewStats() *Stats {
	return &Stats{}
}

func (s Stats) View() string {
	return "Your stats"
}

func (s Stats) Update(msg tea.Msg) (content.Content, tea.Cmd) {
	return s, nil
}
