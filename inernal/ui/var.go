package ui

import (
	"main/inernal/models"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	MainIdx = iota
	ConnectIdx
	SettingsIdx
)

var (
	panelWidth  = 30
	panelHeight = 17

	InfoBar = lipgloss.NewStyle().
		Background(lipgloss.Color("#333333")).
		Foreground(lipgloss.Color("#ffffff")).
		Padding(0, 1)

	// Поле ввода
	inputStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Width(panelWidth - 2).
			Height(1).
			Foreground(lipgloss.Color("#3b6380")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#3b6380"))

	callPanel = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#3e3963")).
			Width(18).
			Height(8).
			Padding(1, 1).
			BorderLeft(false)
)

func MainStyle(w, h int) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#3e3963")). // мягкий голубой
		Width(w).
		Height(h).
		Padding(1, 1)

}

func ConnPanel(m *Model) string {
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#3b6380")).
		Render(m.Connect.Name)
	return MainStyle(panelWidth, panelHeight).Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			title,
			m.Connect.TextInput.View(),
		),
	)
}

func textColor(text, color string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(color)).Render(text)

}

func Calling(name string) string {
	ct := lipgloss.JoinVertical(lipgloss.Center, "Call with", name, lipgloss.JoinHorizontal(lipgloss.Center,
		"end call: ",
		textColor("ctrl+n", "#a01717ff")))

	return callPanel.Render(ct)
}

func GetCall(name string) string {
	return callPanel.Render(lipgloss.JoinVertical(lipgloss.Center, "Call from", name,
		lipgloss.JoinHorizontal(lipgloss.Center,
			textColor("ctrl+y", "#0d881dff"),
			textColor(" // ", "#474747ff"),
			textColor("ctrl+n", "#a01717ff"))))
}

func ToCall(name string) string {
	return callPanel.Render(lipgloss.JoinVertical(lipgloss.Center, "Call to", name,
		lipgloss.JoinHorizontal(lipgloss.Center,
			"Press to cancel:",
			textColor("ctrl+n", "#a01717ff"))))
}

func chatPanel(conn *models.Conn) string {
	conn.ViewPort.Height = (panelHeight - 2) - inputStyle.GetHeight() - 1

	conn.TextInput.Width = panelWidth - 2

	return MainStyle(panelWidth, panelHeight).Render(lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().
			Padding(0, 1).
			Width(panelWidth-2).
			MaxHeight(conn.ViewPort.Height).
			Render(conn.ViewPort.View()),
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("##3e5463")).
			Render(strings.Repeat("-", panelWidth-2)),
		lipgloss.NewStyle().
			MaxHeight(1).
			Render(conn.TextInput.View()),
	))
}

func SelectLang(m *Model) string {
	style := MainStyle(panelWidth, panelHeight)

	return style.Render(lipgloss.JoinVertical(lipgloss.Center, m.Main.LangList.View()))
}

func column(strs ...string) string {
	return lipgloss.JoinVertical(lipgloss.Left, strs...)
}

func (m *Model) ExampleScreen(list, bar string) string {
	return m.vert(lipgloss.JoinHorizontal(
		lipgloss.Top,
		MainStyle(
			panelWidth-10,
			panelHeight-4).
			Render(
				list)),
		InfoBar.Render(m.MakeHelpBar(bar)))
}

func (m *Model) MainScreen(conn *models.Conn) string {
	return m.vert(lipgloss.JoinHorizontal(
		lipgloss.Top,
		MainStyle(
			panelWidth-10,
			panelHeight-4).
			BorderRight(false).
			Render(
				m.Main.ConnList.View()),
		chatPanel(conn)),
		InfoBar.Render(m.MakeHelpBar("chat")))
}
