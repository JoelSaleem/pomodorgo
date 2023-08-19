package timers

import (
	"github.com/JoelSaleem/pomodorgo/internal/content"
	tea "github.com/charmbracelet/bubbletea"
)

type Timers struct{}

func NewTimers() *Timers {
	return &Timers{}
}

func (t Timers) View() string {
	return "Your timers"
}

func (t Timers) Update(msg tea.Msg) (content.Content, tea.Cmd) {
	return nil, nil
}
