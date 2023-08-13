package tabs

import (
	"github.com/JoelSaleem/pomodorgo/internal/stats"
	"github.com/JoelSaleem/pomodorgo/internal/tasks"
	"github.com/JoelSaleem/pomodorgo/internal/timers"
)

type Tab struct {
	Name    string
	content IContent
}

func (t Tab) RenderContent() string {
	return t.content.View()
}

func ConstructTabs() []Tab {
	return []Tab{
		{Name: "Tasks", content: tasks.NewTasks()},
		{Name: "Timers", content: timers.NewTimers()},
		{Name: "Stats", content: stats.NewStats()},
		{Name: "Settings", content: timers.NewTimers()},
	}
}

type IContent interface {
	View() string
}
