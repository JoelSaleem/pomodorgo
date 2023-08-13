package tabs

var Tabs = []Tab{
	{"Tasks", "My Tasks"},
	{"Timers", "Timers go here"},
	{"Stats", "Stats go here"},
	{"Settings", "Settings go here"},
}

type Tab struct {
	Name    string
	content string
}

func (t Tab) RenderContent() string {
	return t.content
}
