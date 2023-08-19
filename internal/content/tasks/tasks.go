package tasks

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// add task (persistent) -- title, description, due date, priority, status
// soft delete task (persistent)
// update task (persistent)
// view tasks
// archive task -- not shown in list
// unarchive task

type listKeyMap struct {
	toggleSpinner    key.Binding
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
	insertItem       key.Binding
}

type Tasks struct {
	list list.Model
	keys *listKeyMap
}

func NewTasks() *Tasks {
	return &Tasks{}
}

func (t Tasks) View() string {
	return "Your tasks"
}

func (t Tasks) Update(msg tea.Msg) (tea.Cmd) {
	return nil
}
