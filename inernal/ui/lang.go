package ui

import (
	"fmt"
	"main/inernal/models"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) SetupLang() {

	m.MainList.SetItems(nil)

	main := []models.Item{
		{Name: m.Language.SelectedLang.ConnectUser},
		{Name: m.Language.SelectedLang.Chats},
		{Name: m.Language.SelectedLang.Settings},
	}
	for idx, item := range main {
		m.MainList.InsertItem(idx, item)
	}

	m.ConnList.SetItems(nil)
	m.ConnList.InsertItem(0, models.Item{Name: m.Language.SelectedLang.BackMain})

	m.SettingsList.SetItems(nil)

	Setttings := []models.Item{
		{Name: m.Language.SelectedLang.BackMain},
		{Name: m.Language.SelectedLang.Lang}}

	for idx, item := range Setttings {
		m.SettingsList.InsertItem(idx, item)
	}
	m.UserConnect.Header = m.Language.SelectedLang.HeaderConnect
	m.TextInput.Placeholder = m.Language.SelectedLang.EntryInput
}

func (m *Model) LangInit() {
	for _, i := range m.Language.Language {
		m.Language.Langs = append(m.Language.Langs, i.Language)
	}
	m.Language.SelectedLang = m.Language.Language[0]

	mainList := MakeList([]models.Item{
		{Name: m.Language.SelectedLang.ConnectUser},
		{Name: m.Language.SelectedLang.Chats},
		{Name: m.Language.SelectedLang.Settings},
	})
	chatList := MakeList([]models.Item{{Name: m.Language.SelectedLang.BackMain}})
	Settings := MakeList([]models.Item{
		{Name: m.Language.SelectedLang.BackMain},
		{Name: m.Language.SelectedLang.Lang},
	})
	notes := []models.Item{}

	for _, i := range m.Language.Langs {
		notes = append(notes, models.Item{Name: i})
	}

	LangList := MakeList(notes)
	LangList.Title = ""
	LangList.Styles.Title = lipgloss.Style{}

	m.MainList = mainList
	m.ConnList = chatList
	m.SettingsList = Settings
	m.LangList = LangList
	m.Screen = MainIdx

	m.UserConnect.Header = m.Language.SelectedLang.HeaderConnect
	m.TextInput = Textinput(m.Language.SelectedLang.EntryInput)

}

func (m *Model) MakeHelpBar(bar string) string {
	switch bar {
	case "main":
		return fmt.Sprintf("%s %s %s", m.Language.SelectedLang.MoveLists, m.Language.SelectedLang.MoveToLists, m.Language.SelectedLang.ExitButton)
	case "chat":
		return fmt.Sprintf("%s %s %s", m.Language.SelectedLang.ScrollMessage, m.Language.SelectedLang.MoveLists, m.Language.SelectedLang.MoveToLists)
	}
	return ""
}

func MakeList(notes []models.Item) list.Model {
	items := make([]list.Item, len(notes))
	for i, n := range notes {
		items[i] = n
	}
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	l := list.New(items, delegate, 20, 15)
	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.Title = "Narria"
	l.SetShowHelp(false)
	return l
}

func Textinput(Placeholder string) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = Placeholder
	return ti
}
