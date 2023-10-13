package add_tasks

import (
	"github.com/JoelSaleem/pomodorgo/internal/content"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AddTasks struct {
	textInput *textinput.Model
}

func (a AddTasks) View() string {
	return a.textInput.View()
}

func (a AddTasks) Update(msg tea.Msg) (content.Content, tea.Cmd) {
	var cmd tea.Cmd
	// a.textInput, cmd = a.textInput.Update(msg)
	return a, cmd
}

func NewAddTasks(width int) *AddTasks {
	ti := textinput.New()
	ti.Placeholder = "Add a task"
	ti.Focus()
	ti.Width = width - 2

	return &AddTasks{
		textInput: &ti,
	}
}
