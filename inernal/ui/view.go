package ui

import (
	"main/inernal/models"

	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	if m.HelloScreen {
		m.TextInput.Placeholder = m.Language.SelectedLang.EntryInput
		return column(textColor(m.Language.SelectedLang.EntryLabel, "#9d9e49ff"), m.TextInput.View())
	}
	if m.Screen == SettingsIdx && m.Language.LangUpd || m.SettingsList.Index() == 1 {
		return m.ExampleScreen(m.SettingsList.View(), m.SelectLang(), "main", false)

	}
	if m.MainList.Index() == 0 {
		return m.ExampleScreen(m.MainList.View(), m.ConnPanel(), "main", false)
	}
	return m.screen()
}

func (m *Model) IfCall(str, Info string) string {
	var call string
	if m.Call.FromCall {
		call = m.GetCall(m.Call.Name)
	}
	if m.Call.ToCall {
		call = m.ToCall(m.Call.Name)
	}
	if m.Call.InCall {
		call = m.Calling(m.Call.Name)
	}
	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.JoinHorizontal(
		lipgloss.Top, str, call), Info)
}

func (m *Model) screen() string {
	switch m.Screen {
	case MainIdx:
		return m.ExampleScreen(m.MainList.View(), "", "main", true)
	case SettingsIdx:
		return m.ExampleScreen(m.SettingsList.View(), "", "chat", true)
	}

	if m.ConnList.Cursor() == 0 {
		return m.ExampleScreen(m.ConnList.View(), "", "chat", true)
	}
	var conn models.Conn
	for _, i := range m.UserConnect.List {
		if i.Id == m.ConnList.Cursor()+1 {
			conn = *i
			break
		}
	}
	if conn.Conn == nil {
		return m.ExampleScreen(m.ConnList.View(), "", "chat", true)
	}

	return m.MainScreen(&conn)
}
