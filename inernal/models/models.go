package models

import (
	"main/inernal/localization"
	"net"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
)

// Сообщения из пакетов в update bubbletea
type ServerMsg struct {
	Conn    net.Conn
	Service ServiceMsg
	Msg     string
}

type ServiceMsg struct {
	Type string
	Msg  string
}

type UserMessage struct {
	Message     string
	IsHandShake bool
}

func (s *ServiceMsg) SetValue(msg, service string) {
	s.Msg = msg
	s.Type = service
}

// Структура звонка
type Call struct {
	InCall   bool
	FromCall bool
	ToCall   bool
	Conn     *net.Conn
	Name     string
	Timer    int
	TimerOn  chan struct{}
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

type UserConnect struct {
	Header    string
	TextInput textinput.Model
	IsEditing bool
	Name      string
	List      []*Conn
}

type Lang struct {
	Language     []localization.Lang
	Langs        []string
	SelectedLang localization.Lang
	LangIdx      int
	LangUpd      bool
}

type Item struct {
	Name string
}

func (n Item) Title() string       { return n.Name }
func (n Item) Description() string { return "" }
func (n Item) FilterValue() string { return n.Name }
