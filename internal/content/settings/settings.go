package settings

import tea "github.com/charmbracelet/bubbletea"

type Settings struct{}

func NewSettings() *Settings {
	return &Settings{}
}

func (t Settings) View() string {
	return "Your settings"
}

func (s Settings) Update(msg tea.Msg) tea.Cmd {
	return nil
}
