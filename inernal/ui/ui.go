package ui

import (
	"main/inernal/models"
	"main/inernal/peer"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Username *string
	Program  *tea.Program

	MainList     list.Model
	ConnList     list.Model
	SettingsList list.Model
	LangList     list.Model

	Screen uint

	TextInput textinput.Model

	HelloScreen bool

	//Для настройки и обработки языков используеться модешь models.lang
	Language models.Lang

	//Подключение к пользователям и передача сообщений как пользовательских
	// так и системных используеться tcp подключение реализованное в peer.Peer
	Peer peer.Peer

	//Здесь храняться все чаты и основные компоненты окна подключения
	UserConnect models.UserConnect

	//структура звонков
	Call models.Call
}

func (m *Model) Init() tea.Cmd { return nil }
