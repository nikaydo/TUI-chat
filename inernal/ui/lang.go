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
		{Name: m.SelectedLang.ConnectUser},
		{Name: m.SelectedLang.Chats},
		{Name: m.SelectedLang.Settings},
	}
	for idx, item := range main {
		m.MainList.InsertItem(idx, item)
	}

	m.ConnList.SetItems(nil)
	m.ConnList.InsertItem(0, models.Item{Name: m.SelectedLang.BackMain})

	m.SettingsList.SetItems(nil)

	Setttings := []models.Item{
		{Name: m.SelectedLang.BackMain},
		{Name: m.SelectedLang.Lang}}

	for idx, item := range Setttings {
		m.SettingsList.InsertItem(idx, item)
	}

	m.Hello.TextInput.Placeholder = m.SelectedLang.EntryInput
}

func (m *Model) LangInit() {
	for _, i := range m.Language {
		m.Langs = append(m.Langs, i.Language)
	}
	m.SelectedLang = m.Language[0]

	mainList := MakeList([]models.Item{
		{Name: m.SelectedLang.ConnectUser},
		{Name: m.SelectedLang.Chats},
		{Name: m.SelectedLang.Settings},
	})
	chatList := MakeList([]models.Item{{Name: m.SelectedLang.BackMain}})
	Settings := MakeList([]models.Item{
		{Name: m.SelectedLang.BackMain},
		{Name: m.SelectedLang.Lang},
	})
	notes := []models.Item{}

	for _, i := range m.Langs {
		notes = append(notes, models.Item{Name: i})
	}

	LangList := MakeList(notes)
	LangList.Title = ""
	LangList.Styles.Title = lipgloss.Style{}
	m.Main = models.Main{
		MainList:     mainList,
		ConnList:     chatList,
		SettingsList: Settings,
		LangList:     LangList,
		Screen:       MainIdx,
		Err:          false}

	m.Hello.TextInput = Textinput(m.SelectedLang.EntryInput)

}

func (m *Model) MakeHelpBar(bar string) string {
	switch bar {
	case "main":
		return fmt.Sprintf("%s %s %s", m.SelectedLang.MoveLists, m.SelectedLang.MoveToLists, m.SelectedLang.ExitButton)
	case "chat":
		return fmt.Sprintf("%s %s %s", m.SelectedLang.ScrollMessage, m.SelectedLang.MoveLists, m.SelectedLang.MoveToLists)
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
