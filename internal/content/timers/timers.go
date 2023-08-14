package timers

type Timers struct {}

func NewTimers() *Timers {
	return &Timers{}
}

func (t Timers) View() string {
	return "Your timers"
}