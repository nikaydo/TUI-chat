package ui

import (
	"main/inernal/models"

	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	if m.Hello.IsEditing {
		m.Hello.TextInput.Placeholder = m.SelectedLang.EntryInput
		return column(textColor(m.SelectedLang.EntryLabel, "#9d9e49ff"), m.Hello.TextInput.View())
	}
	if m.Screen == SettingsIdx {
		if m.LangUpd {

			return m.vert(lipgloss.JoinHorizontal(
				lipgloss.Top,
				MainStyle(panelWidth-10, panelHeight-4).BorderRight(false).Render(m.Main.SettingsList.View()),
				SelectLang(m)), InfoBar.Render(m.MakeHelpBar("main")))
		}
		switch m.Main.SettingsList.Index() {
		case 1:
			return m.vert(lipgloss.JoinHorizontal(
				lipgloss.Top,
				MainStyle(panelWidth-10, panelHeight-4).BorderRight(false).Render(m.Main.SettingsList.View()),
				SelectLang(m)), InfoBar.Render(m.MakeHelpBar("main")))
		}
	}
	if m.Main.MainList.Index() == 0 {
		return m.vert(lipgloss.JoinHorizontal(
			lipgloss.Top,
			MainStyle(panelWidth-10, panelHeight-4).BorderRight(false).Render(m.Main.MainList.View()),
			ConnPanel(m)), InfoBar.Render(m.MakeHelpBar("main")))
	}
	return screen(m)
}

func (m *Model) vert(str, Info string) string {
	var call string
	if m.Call.FromCall {
		call = GetCall(m.Call.Name)
	}
	if m.Call.ToCall {
		call = ToCall(m.Call.Name)
	}
	if m.Call.InCall {
		call = Calling(m.Call.Name)
	}
	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.JoinHorizontal(
		lipgloss.Top, str, call), Info)
}

func screen(m *Model) string {
	switch m.Screen {
	case MainIdx:
		return m.ExampleScreen(m.Main.MainList.View(), "main")
	case SettingsIdx:
		return m.ExampleScreen(m.Main.SettingsList.View(), "chat")

	}

	if m.Main.ConnList.Cursor() == 0 {
		return m.ExampleScreen(m.Main.ConnList.View(), "chat")
	}
	var conn models.Conn
	for _, i := range m.Connect.List {
		if i.Id == m.Main.ConnList.Cursor()+1 {
			conn = *i
			break
		}
	}
	if conn.Conn == nil {
		return m.ExampleScreen(m.Main.ConnList.View(), "chat")
	}

	return m.MainScreen(&conn)
}
