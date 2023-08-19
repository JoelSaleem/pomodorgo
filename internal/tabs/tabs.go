package tabs

import (
	"github.com/JoelSaleem/pomodorgo/internal/content"
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

func ConstructTabs() []Tab {
	return []Tab{
		{Name: "Tasks", content: tasks.NewTasks()},
		{Name: "Timers", content: timers.NewTimers()},
		{Name: "Stats", content: stats.NewStats()},
		{Name: "Settings", content: timers.NewTimers()},
	}
}
