package models

import (
	"main/inernal/localization"
	"net"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
)

type UserMessage struct {
	Message     string
	IsHandShake bool
	IsCall      bool
	//0 - Call
	//1 - Accept
	//2 - End call or decline
	CallStatus int
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
