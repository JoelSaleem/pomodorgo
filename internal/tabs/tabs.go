package tabs

import (
	"github.com/JoelSaleem/pomodorgo/internal/content"
	"github.com/JoelSaleem/pomodorgo/internal/content/add_tasks"
	"github.com/JoelSaleem/pomodorgo/internal/content/stats"
	"github.com/JoelSaleem/pomodorgo/internal/content/tasks"
	"github.com/JoelSaleem/pomodorgo/internal/content/timers"
	tea "github.com/charmbracelet/bubbletea"
)

type Tab struct {
	Name    string
	content content.Content
}

func (t Tab) RenderContent() string {
	return t.content.View()
}

func (t Tab) Update(msg tea.Msg) (Tab, tea.Cmd) {
	content, cmd := t.content.Update(msg)
	t.content = content
	return t, cmd
}

func ConstructTabs(width, height int) []Tab {
	return []Tab{
		{Name: "Add a task", content: add_tasks.NewAddTasks(width)},
		{Name: "My tasks", content: tasks.NewTasks(width, height)},
		{Name: "Timers", content: timers.NewTimers()},
		{Name: "Stats", content: stats.NewStats()},
		{Name: "Settings", content: timers.NewTimers()},
	}
}
