package stats

type Stats struct {}

func NewStats() *Stats {
	return &Stats{}
}

func (t Stats) View() string {
	return "Your stats"
}