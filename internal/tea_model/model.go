package tea_model

import (
	"github.com/JoelSaleem/pomodorgo/internal/db"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	repository db.IRepository // TODO: must extend interface

	altscreen bool
	quitting  bool

}

func NewProgram(repo *db.Repository) *tea.Program {
	return tea.NewProgram(model{
		repository: repo,
	})
}

func (m model) Init() tea.Cmd {
	m.altscreen = false
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.altscreen {
		m.altscreen = true
		return m, tea.EnterAltScreen
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case " ":

		}
	}
	return m, nil
}

func (m model) View() string {
	tea.EnterAltScreen()
	if m.quitting {
		return "Bye!\n"
	}

	return "press ctrl+c to quit\n"
}
