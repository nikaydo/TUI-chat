package ui

import (
	"main/inernal/models"
	"net"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func check(m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.Screen == ConnectIdx {
		if conn := getConn(m); conn != nil {
			conn.UnReadMsg = 0
			m.Main.ConnList.SetItem(conn.Id-1, models.Item{Name: conn.UserName})
			conn.TextInput.Focus()
			conn.TextInput, cmd = conn.TextInput.Update(msg)
		}
		return m, cmd
	}

	if m.Hello.IsEditing {
		m.Hello.TextInput.Focus()
		m.Hello.TextInput, cmd = m.Hello.TextInput.Update(msg)
		return m, cmd
	}

	if m.Main.MainList.Cursor() == 0 {
		m.Connect.TextInput.Focus()
		m.Connect.TextInput, cmd = m.Connect.TextInput.Update(msg)
		return m, cmd
	}
	return m, cmd
}

func getConn(m *Model) *models.Conn {
	var c *models.Conn
	for _, i := range m.Connect.List {
		if i.Id == m.Main.ConnList.Cursor()+1 {
			return i
		}
	}

	return c
}

func (m *Model) getConnByAddr(addr net.Addr) *models.Conn {
	for _, i := range m.Connect.List {
		c := *i.Conn
		if c.RemoteAddr() == addr {
			return i
		}
	}
	return nil
}

func UserInit(c *net.Conn, m *Model) {
	var conn models.Conn
	conn.Ip = m.Connect.TextInput.Value()
	conn.Conn = c
	conn.ViewPort = viewport.New(50, 13)
	conn.ViewPort.Width = 28
	conn.TextInput = textinput.New()
	conn.Id = len(m.Main.ConnList.Items()) + 1
	m.Connect.List = append(m.Connect.List, &conn)
}

func ConnectUser(m *Model) {
	msg := models.ServerMsg{}
	c, err := m.Peer.Tcp.Connect(m.Connect.TextInput.Value())
	if err != nil {
		if strings.Contains(err.Error(), "i/o timeout") {
			msg.Text.SetValue(m.SelectedLang.ConnectTimeout, "ConnTimeout")
			m.Program.Send(msg)
			return
		}
		msg.Text.SetValue(m.SelectedLang.ConnectError, "ConnTimeout")
		m.Program.Send(msg)
		return
	}
	UserInit(c, m)

	m.Peer.Tcp.SendMsg(c, *m.Username, "HandServer")

	msg.Text.SetValue(textColor(m.SelectedLang.ConnectSucessful, "#0d881dff"), "ConnTimeout")
	m.Program.Send(msg)
}

func move(m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.Screen {
	case MainIdx:
		m.Main.MainList, cmd = m.Main.MainList.Update(msg)
		return m, cmd
	case ConnectIdx:
		m.Main.ConnList, cmd = m.Main.ConnList.Update(msg)
		return m, cmd
	case SettingsIdx:
		m.Main.SettingsList, cmd = m.Main.SettingsList.Update(msg)
		return m, cmd

	}
	return m, cmd
}
