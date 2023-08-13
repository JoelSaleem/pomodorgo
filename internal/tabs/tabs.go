package tabs

var Tabs = []Tab{
	{Name: "Tasks", Content: "My Tasks"},
	{Name: "Timers", Content: "Timers go here"},
	{Name: "Stats", Content: "Stats go here"},
	{Name: "Settings", Content: "Settings go here"},
}

type Tab struct {
	Name    string
	Content string
}

func (t Tab) RenderContent() string {
	return t.Content
}
