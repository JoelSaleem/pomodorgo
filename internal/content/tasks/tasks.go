package tasks

import (
	"github.com/JoelSaleem/pomodorgo/internal/content"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// add task (persistent) -- title, description, due date, priority, status
// soft delete task (persistent)
// update task (persistent)
// view tasks
// archive task -- not shown in list
// unarchive task

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

type listKeyMap struct {
	toggleSpinner    key.Binding
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
	insertItem       key.Binding
}

type Tasks struct {
	list          list.Model
	keys          *listKeyMap
	text          string
	height, width int
}

func (t Tasks) SetDimensions(width, height int) Tasks {
	t.height = height
	t.width = width
	return t
}

type listItem struct {
	title string
}

func (l listItem) FilterValue() string {
	return l.title
}
func (l listItem) Title() string {
	return l.title
}
func (l listItem) Description() string {
	return l.title
}

func NewTasks(width, height int) *Tasks {
	items := make([]list.Item, 0)

	items = append(items, listItem{title: "Task 1"}, listItem{title: "Task 2"}, listItem{title: "Task 3"})
	d := list.NewDefaultDelegate()
	l := list.New(items, d, width, height)
	l.Title = "My Tasks"
	l.Styles.Title = titleStyle

	t := Tasks{
		list: l,
	}.SetDimensions(width, height)

	return &t
}

func (t Tasks) View() string {
	return appStyle.Render(t.list.View())
}

func (t Tasks) Update(msg tea.Msg) (content.Content, tea.Cmd) {
	newListModel, cmd := t.list.Update(msg)
	t.list = newListModel
	return t, cmd
}
