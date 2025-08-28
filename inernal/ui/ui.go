package ui

import (
	"main/inernal/localization"
	"main/inernal/models"
	"main/inernal/peer"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Username     *string
	Program      *tea.Program
	Language     []localization.Lang
	Langs        []string
	SelectedLang localization.Lang
	LangIdx      int
	LangUpd      bool
	//Структура с окном приветствия
	peer.Peer
	models.Hello
	models.Main
	models.Connect
	models.Call
}

func (m *Model) Init() tea.Cmd { return nil }
