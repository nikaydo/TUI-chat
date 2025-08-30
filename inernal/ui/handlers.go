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
			m.ConnList.SetItem(conn.Id-1, models.Item{Name: conn.UserName})
			conn.TextInput.Focus()
			conn.TextInput, cmd = conn.TextInput.Update(msg)
		}
		return m, cmd
	}

	if m.HelloScreen {
		m.TextInput.Focus()
		m.TextInput, cmd = m.TextInput.Update(msg)
		return m, cmd
	}

	if m.MainList.Cursor() == 0 {
		m.UserConnect.TextInput.Focus()
		m.UserConnect.TextInput, cmd = m.UserConnect.TextInput.Update(msg)
		return m, cmd
	}
	return m, cmd
}

func getConn(m *Model) *models.Conn {
	var c *models.Conn
	for _, i := range m.UserConnect.List {
		if i.Id == m.ConnList.Cursor()+1 {
			return i
		}
	}

	return c
}

func (m *Model) getConnByAddr(addr net.Addr) *models.Conn {
	for _, i := range m.UserConnect.List {
		c := *i.Conn
		if c.RemoteAddr() == addr {
			return i
		}
	}
	return nil
}

func UserInit(c *net.Conn, m *Model) {
	var conn models.Conn
	conn.Ip = m.UserConnect.TextInput.Value()
	conn.Conn = c
	conn.ViewPort = viewport.New(50, 13)
	conn.ViewPort.Width = 28
	conn.TextInput = textinput.New()
	conn.Id = len(m.ConnList.Items()) + 1
	m.UserConnect.List = append(m.UserConnect.List, &conn)
}

func ConnectUser(m *Model) {
	c, err := m.Peer.Tcp.Connect(m.UserConnect.TextInput.Value())
	if err != nil {
		if strings.Contains(err.Error(), "i/o timeout") {
			msg := models.ConnTimeout{Data: m.Language.SelectedLang.ConnectTimeout}
			m.Program.Send(msg)
			return
		}
		msg := models.ConnTimeout{Data: m.Language.SelectedLang.ConnectError}
		m.Program.Send(msg)
		return
	}
	UserInit(c, m)

	m.Peer.Tcp.SendMsg(c, models.UserMessage{Message: *m.Username, IsHandShake: true})
	msg := models.ConnTimeout{Data: textColor(m.Language.SelectedLang.ConnectSucessful, "#0d881dff")}
	m.Program.Send(msg)
}

func move(m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.Screen {
	case MainIdx:
		m.MainList, cmd = m.MainList.Update(msg)
		return m, cmd
	case ConnectIdx:
		m.ConnList, cmd = m.ConnList.Update(msg)
		return m, cmd
	case SettingsIdx:
		m.SettingsList, cmd = m.SettingsList.Update(msg)
		return m, cmd

	}
	return m, cmd
}

/*
func (m *Model) TimerCount(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
		default:
			//	s := models.ServerMsg{Service.Type: ""}
		}
	}
}
*/
