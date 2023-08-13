package settings

type Settings struct {}

func NewSettings() *Settings {
	return &Settings{}
}

func (t Settings) View() string {
	return "Your settings"
}