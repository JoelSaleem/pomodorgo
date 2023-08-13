package tea_model

import (
	"fmt"
	"strings"

	"github.com/JoelSaleem/pomodorgo/internal/db"
	"github.com/JoelSaleem/pomodorgo/internal/tabs"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

type model struct {
	repository db.IRepository // TODO: must extend interface

	altscreen bool
	quitting  bool

	activeTab int
}

func NewProgram(repo *db.Repository) *tea.Program {
	return tea.NewProgram(model{
		repository: repo,
	})
}

func (m model) Init() tea.Cmd {
	m.altscreen = false
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.altscreen {
		m.altscreen = true
		fmt.Println("entering alt screen")
		return m, tea.EnterAltScreen
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case " ":
			return m, nil
		case "right", "l", "n", "tab":
			m.activeTab = min(m.activeTab+1, len(tabs.Tabs)-1)
			return m, nil
		case "left", "h", "p", "shift+tab":
			m.activeTab = max(m.activeTab-1, 0)
			return m, nil

		}
	}
	return m, nil
}

func (m model) View() string {
	doc := strings.Builder{}
	tea.EnterAltScreen()
	if m.quitting {
		return "Bye!\n"
	}

	var renderedTabs []string

	for i, t := range tabs.Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(tabs.Tabs)-1, i == m.activeTab
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t.Name))
	}

	doc.WriteString("Pomodorgo\n\n")
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(tabs.Tabs[m.activeTab].RenderContent()))
	return docStyle.Render(doc.String())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
