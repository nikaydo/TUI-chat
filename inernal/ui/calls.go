package ui

import (
	"main/inernal/models"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) chatCalls(key string) tea.Model {
	switch key {
	case "ctrl+f":
		if !m.Call.InCall && !m.Call.FromCall && !m.Call.ToCall {
			if conn := getConn(m); conn != nil {
				m.Peer.Tcp.SendMsg(conn.Conn, "call", "call")
				changeCall(m, "to")
				m.Call.Conn = conn.Conn
				m.Call.Name = conn.UserName
			}
		}
		return m
	case "ctrl+y":
		if m.Call.FromCall {
			m.Peer.Tcp.SendMsg(m.Call.Conn, "accept", "call")

			changeCall(m, "in")
		}
		return m
	case "ctrl+n":
		m.Peer.Tcp.SendMsg(m.Call.Conn, "decline", "call")
		changeCall(m, "")
		m.Call.Name = ""
	}
	return m
}

func (m *Model) commandCalls(msg models.ServerMsg) tea.Model {
	switch msg.Text.Msg {
	case "call":
		if User := m.getConnByAddr(msg.Conn.RemoteAddr()); User != nil {
			m = changeCall(m, "get")
			m.Call.Name = User.UserName
			m.Call.Conn = User.Conn
		}
		return m
	case "accept":
		if User := m.getConnByAddr(msg.Conn.RemoteAddr()); User != nil {
			m = changeCall(m, "in")
			m.Call.Name = User.UserName
			m.Call.Conn = User.Conn
		}
		return m
	default:
		if User := m.getConnByAddr(msg.Conn.RemoteAddr()); User != nil {
			m = changeCall(m, "")
			m.Call.Name = User.UserName
			m.Call.Conn = nil
		}
		return m
	}
}

func changeCall(m *Model, t string) *Model {
	switch t {
	case "in":
		m.Call.FromCall = false
		m.Call.ToCall = false
		m.Call.InCall = true
	case "to":
		m.Call.FromCall = false
		m.Call.ToCall = true
		m.Call.InCall = false
	case "get":
		m.Call.FromCall = true
		m.Call.ToCall = false
		m.Call.InCall = false
	default:
		m.Call.FromCall = false
		m.Call.ToCall = false
		m.Call.InCall = false
	}
	return m
}
