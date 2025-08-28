package models

import (
	"net"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
)

type ServerMsg struct {
	Conn net.Conn
	Text ServiceMsg
}

type ServiceMsg struct {
	Service string
	Msg     string
}

func (s *ServiceMsg) SetValue(msg, service string) {
	s.Msg = msg
	s.Service = service
}

type Call struct {
	InCall   bool
	FromCall bool
	ToCall   bool
	Conn     *net.Conn
	Name     string
}

type Connections struct {
	List []*Conn
}

type Conn struct {
	Id        int
	Ip        string
	UserName  string
	Conn      *net.Conn
	ViewPort  viewport.Model
	TextInput textinput.Model
	Msg       []string
	UnReadMsg uint
}

type Hello struct {
	TextInput textinput.Model
	IsEditing bool
}

type Main struct {
	MainList     list.Model
	ConnList     list.Model
	SettingsList list.Model

	LangList list.Model
	Screen   uint
	Err      bool
}

type Connect struct {
	Header    string
	TextInput textinput.Model
	IsEditing bool
	Name      string
	List      []*Conn
}

type Item struct {
	Name string
}

func (n Item) Title() string       { return n.Name }
func (n Item) Description() string { return "" }
func (n Item) FilterValue() string { return n.Name }
