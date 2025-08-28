package main

import (
	"fmt"
	"log"
	"main/inernal/config"
	localization "main/inernal/localization"
	"main/inernal/models"
	"main/inernal/peer"
	"main/inernal/ui"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cfg, err := config.ReadEnv()
	if err != nil {
		log.Fatalln(err)
	}
	var m ui.Model
	p := tea.NewProgram(&m)
	lang, err := localization.LangRead(cfg.LangPath)
	if err != nil {
		log.Fatalln(err)
	}

	m = ui.Model{
		LangIdx:  0,
		Language: lang,
		Hello:    models.Hello{IsEditing: true},
		Connect:  models.Connect{TextInput: textinput.New(), IsEditing: false},
		Program:  p,
		Peer: peer.Peer{
			Tcp: peer.Tcp{
				Port:     cfg.TCP.Port,
				Host:     cfg.TCP.Host,
				Programm: p}},
	}
	m.LangInit()
	if len(os.Args) != 1 {
		m.Tcp.RunServers()
	}
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
