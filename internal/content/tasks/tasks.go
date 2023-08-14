package tasks

// add task (persistent) -- title, description, due date, priority, status
// soft delete task (persistent)
// update task (persistent)
// view tasks
// archive task -- not shown in list
// unarchive task
type Tasks struct {}

func NewTasks() *Tasks {
	return &Tasks{}
}

func (t Tasks) View() string {
	return "Your tasks"
}