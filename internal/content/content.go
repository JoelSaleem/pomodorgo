package content

import tea "github.com/charmbracelet/bubbletea"

type Content interface {
	View() string
	Update(msg tea.Msg) (Content, tea.Cmd)
}
