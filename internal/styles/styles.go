package styles

import "github.com/charmbracelet/lipgloss"

var (
	InactiveTabBorder = TabBorderWithBottom("┴", "─", "┴")
	ActiveTabBorder   = TabBorderWithBottom("┘", " ", "└")
	DocStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	HighlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	InactiveTabStyle  = lipgloss.NewStyle().Border(InactiveTabBorder, true).BorderForeground(HighlightColor).Padding(0, 1)
	ActiveTabStyle    = InactiveTabStyle.Copy().Border(InactiveTabBorder, true)
	WindowStyle       = lipgloss.NewStyle().BorderForeground(HighlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

func TabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}