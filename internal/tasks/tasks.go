package tasks

type Tasks struct {}

func NewTasks() *Tasks {
	return &Tasks{}
}

func (t Tasks) View() string {
	return "Your tasks"
}