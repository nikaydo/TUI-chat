package ui

import (
	"fmt"

	"main/inernal/models"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case models.ServerMsg:
		switch msg.Text.Service {
		case "HandServer":
			if UserConnection := m.getConnByAddr(msg.Conn.RemoteAddr()); UserConnection != nil {
				UserConnection.UserName = msg.Text.Msg
				cmd = m.Main.ConnList.InsertItem(UserConnection.Id, models.Item{Name: UserConnection.UserName})
				return m, cmd
			}
		case "PrepareUser":
			UserInit(&msg.Conn, m)
			return m, cmd
		case "ConnTimeout":
			m.Connect.Name = textColor(msg.Text.Msg, "#b82424ff")
			return m, nil
		case "call":
			md := m.commandCalls(msg)
			return md, nil
		default:
			if UserConnection := m.getConnByAddr(msg.Conn.RemoteAddr()); UserConnection != nil {
				UserConnection.UnReadMsg += 1
				wrappedMsg := wrapMessage(splitText(UserConnection.UserName, msg.Text.Msg, UserConnection.ViewPort.Width), UserConnection.ViewPort.Width)
				UserConnection.Msg = append(UserConnection.Msg, textColor(wrappedMsg, "#ac7cb9ff"))
				UserConnection.ViewPort.SetContent(strings.Join(UserConnection.Msg, "\n"))
				UserConnection.ViewPort.GotoBottom()
				if m.Main.ConnList.Cursor() != UserConnection.Id-1 {
					cmd = m.Main.ConnList.SetItem(UserConnection.Id-1, models.Item{Name: fmt.Sprintf("%s [%d]", UserConnection.UserName, UserConnection.UnReadMsg)})
					return m, cmd
				}
				UserConnection.UnReadMsg = 0
				cmd = m.Main.ConnList.SetItem(UserConnection.Id-1, models.Item{Name: UserConnection.UserName})
			}
			return m, cmd
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			switch m.Screen {
			case 0:
				//show main page
				if m.MainList.Cursor() == 1 {
					m.Screen = ConnectIdx
				}
				if m.MainList.Cursor() == 2 {
					m.Screen = SettingsIdx
				}
				return m, nil
				//show connections page
			case 1:
				if m.ConnList.Cursor() == 0 {
					m.Screen = MainIdx
				}
				return m, nil
			case 2:
				if m.SettingsList.Cursor() == 0 {
					m.Screen = MainIdx
				}
				return m, nil
			}
			//accept input from user
		case "enter":
			//Get user name from textinput
			if m.Hello.IsEditing {
				n := m.Hello.TextInput.Value()
				m.Username = &n
				m.Tcp.Name = &n
				m.Hello.TextInput.Reset()
				m.Hello.IsEditing = false
				return m, nil
			}
			//connect to user
			if m.Main.MainList.Cursor() == 0 {
				m.Connect.Name = textColor(m.SelectedLang.Connectig, "#b18f30ff")
				go ConnectUser(m)
				return m, nil
			}
			//checking if addr textinput is empty to do nothing and if he have addres send message to user in connect page
			if conn := getConn(m); conn != nil {
				if conn.TextInput.Value() == "" {
					return m, cmd
				}
				wrappedMsg := wrapMessage(splitText(*m.Username, conn.TextInput.Value(), conn.ViewPort.Width), conn.ViewPort.Width)
				conn.Msg = append(conn.Msg, textColor(wrappedMsg, "#9e6354ff"))
				conn.ViewPort.SetContent(strings.Join(conn.Msg, "\n"))
				m.Peer.Tcp.SendMsg(conn.Conn, conn.TextInput.Value(), "")
				conn.TextInput.Reset()
				conn.ViewPort.GotoBottom()
			}
			return m, cmd
			//buttons controls calls
		case "ctrl+f", "ctrl+y", "ctrl+n":
			md := m.chatCalls(msg.String())
			return md, nil
			//move in list of main menu and and another lists
		case "up", "down":
			move(m, msg)
			if conn := getConn(m); conn != nil {
				conn.UnReadMsg = 0
				m.Main.ConnList.SetItem(conn.Id-1, models.Item{Name: conn.UserName})
				conn.ViewPort.SetContent(strings.Join(conn.Msg, "\n"))
			}
			return m, nil
			//scroll up message list
		case tea.KeyCtrlW.String():
			switch m.Screen {
			case ConnectIdx:
				if conn := getConn(m); conn != nil {
					conn.ViewPort.ScrollUp(1)
				}
			case SettingsIdx:
				m.langScroll(m.LangIdx-1, false)
				return m, nil
			}
			//scroll down message list
		case "ctrl+s":
			switch m.Screen {
			case ConnectIdx:
				if conn := getConn(m); conn != nil {
					conn.ViewPort.ScrollDown(1)
				}
			case SettingsIdx:
				m.langScroll(m.LangIdx+1, true)
			}
			return m, nil
			//close app
		case "ctrl+c":
			return m, tea.Quit

		}
	}
	return check(m, msg)
}

func (m *Model) langScroll(idx int, move bool) {
	if m.SettingsList.Index() == 1 {
		if move {
			m.Main.LangList.CursorDown()
		} else {
			m.Main.LangList.CursorUp()
		}
		n := m.LangList.SelectedItem().FilterValue()
		for _, i := range m.Language {
			if i.Language == n {
				if idx >= len(m.LangList.Items()) {
					return
				}
				if idx <= -1 {
					return
				}
				m.LangIdx = idx
				m.SelectedLang = i
				m.LangUpd = true
				m.SetupLang()
				m.LangList.Select(m.LangIdx)
				return
			}
		}
	}
}
