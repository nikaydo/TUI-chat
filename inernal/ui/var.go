package ui

import (
	"main/inernal/models"
	"strconv"
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
		BorderForeground(lipgloss.Color("#3e3963")).
		Width(w).
		Height(h).
		Padding(1, 1)

}

func (m *Model) ConnPanel() string {
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#3b6380")).
		Render(m.UserConnect.Header)
	return MainStyle(panelWidth, panelHeight).Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			title,
			m.UserConnect.TextInput.View(),
		),
	)
}

func textColor(text, color string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(color)).Render(text)

}

func (m *Model) Calling(name string) string {

	ct := lipgloss.JoinVertical(lipgloss.Center, m.Language.SelectedLang.CallWith, name, strconv.Itoa(m.Call.Timer), lipgloss.JoinHorizontal(lipgloss.Center,
		m.Language.SelectedLang.CallEnd,
		textColor("ctrl+n", "#a01717ff")))
	return callPanel.Render(ct)
}

func (m *Model) GetCall(name string) string {
	return callPanel.Render(lipgloss.JoinVertical(lipgloss.Center, m.Language.SelectedLang.CallFrom, name,
		lipgloss.JoinHorizontal(lipgloss.Center,
			textColor("ctrl+y", "#0d881dff"),
			textColor(" // ", "#474747ff"),
			textColor("ctrl+n", "#a01717ff"))))
}

func (m *Model) ToCall(name string) string {
	return callPanel.Render(lipgloss.JoinVertical(lipgloss.Center, m.Language.SelectedLang.CallTo, name,
		lipgloss.JoinHorizontal(lipgloss.Center,
			m.Language.SelectedLang.CallCancel,
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

func (m *Model) SelectLang() string {
	style := MainStyle(panelWidth, panelHeight)
	return style.Render(lipgloss.JoinVertical(lipgloss.Center, m.LangList.View()))
}

func column(strs ...string) string {
	return lipgloss.JoinVertical(lipgloss.Left, strs...)
}

func (m *Model) ExampleScreen(list, left, bar string, right bool) string {
	main := MainStyle(
		panelWidth-10,
		panelHeight-4,
	).BorderRight(right)

	return m.IfCall(lipgloss.JoinHorizontal(
		lipgloss.Top,
		main.Render(list),
		left),
		InfoBar.Render(m.MakeHelpBar(bar)))

}

func (m *Model) MainScreen(conn *models.Conn) string {
	return m.IfCall(lipgloss.JoinHorizontal(
		lipgloss.Top,
		MainStyle(
			panelWidth-10,
			panelHeight-4).
			BorderRight(false).
			Render(
				m.ConnList.View()),
		chatPanel(conn)),
		InfoBar.Render(m.MakeHelpBar("chat")))
}

func (m *Model) langScroll(idx int, move bool) {
	if m.SettingsList.Index() == 1 {
		if move {
			m.LangList.CursorDown()
		} else {
			m.LangList.CursorUp()
		}
		n := m.LangList.SelectedItem().FilterValue()
		for _, i := range m.Language.Language {
			if i.Language == n {
				if idx >= len(m.LangList.Items()) {
					return
				}
				if idx <= -1 {
					return
				}
				m.Language.LangIdx = idx
				m.Language.SelectedLang = i
				m.Language.LangUpd = true
				m.SetupLang()
				m.LangList.Select(m.Language.LangIdx)
				return
			}
		}
	}
}
