package tabs

import (
	"github.com/JoelSaleem/pomodorgo/internal/content"
	"github.com/JoelSaleem/pomodorgo/internal/content/stats"
	"github.com/JoelSaleem/pomodorgo/internal/content/tasks"
	"github.com/JoelSaleem/pomodorgo/internal/content/timers"
)

type Tab struct {
	Name    string
	content content.Content
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
