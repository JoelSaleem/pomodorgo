package tea_model

import (
	"strings"

	"github.com/JoelSaleem/pomodorgo/internal/db"
	"github.com/JoelSaleem/pomodorgo/internal/styles"
	"github.com/JoelSaleem/pomodorgo/internal/tabs"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	repository db.IRepository // TODO: must extend interface
	quitting   bool
	activeTab  int
	tabs       []tabs.Tab
}

func NewProgram(repo *db.Repository, tabs []tabs.Tab) *tea.Program {
	return tea.NewProgram(model{
		repository: repo,
		tabs:       tabs,
	})
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case "right":
			m.activeTab = min(m.activeTab+1, len(m.tabs)-1)
			return m, nil
		case "left":
			m.activeTab = max(m.activeTab-1, 0)
			return m, nil
		// case " ":
		// 	return m, nil
		default:
			tab, cmd := m.tabs[m.activeTab].Update(msg)
			m.tabs[m.activeTab] = tab
			return m, cmd
		}
	}
	return m, nil
}

func (m model) View() string {
	doc := strings.Builder{}
	if m.quitting {
		return "Bye!\n"
	}

	var renderedTabs []string

	for i, t := range m.tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.tabs)-1, i == m.activeTab
		if isActive {
			style = styles.ActiveTabStyle.Copy()
		} else {
			style = styles.InactiveTabStyle.Copy()
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

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString(
		styles.WindowStyle.Height(
			lipgloss.Height(row) - styles.WindowStyle.GetVerticalFrameSize(),
		).Width(
			lipgloss.Width(row) - styles.WindowStyle.GetHorizontalFrameSize(),
		).Render(
			m.tabs[m.activeTab].RenderContent(),
		),
	)
	doc.WriteString("\n")
	return styles.DocStyle.Render(doc.String())
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
