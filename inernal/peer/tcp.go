package peer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"time"

	"main/inernal/models"

	"net"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Tcp struct {
	Port     int
	Host     string
	Programm *tea.Program

	Name *string
}

//Cliet

func (t *Tcp) Connect(addr string) (*net.Conn, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(addr, fmt.Sprintf("%d", t.Port)), time.Second*5)
	if err != nil {
		return &conn, err
	}

	f := func(conn *net.Conn, msg models.ServerMsg) {
		t.Programm.Send(msg)
	}

	go t.handleRequest(&conn, f)
	return &conn, err
}

func (t *Tcp) SendMsg(c *net.Conn, text, service string) {
	conn := *c
	msg := models.ServiceMsg{Type: service, Msg: text}
	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("error marshaling message", err.Error())
		return
	}
	conn.Write(append(b, '\n'))
}

//Server

func (t *Tcp) RunServers() {
	listener, err := net.Listen("tcp", net.JoinHostPort(t.Host, fmt.Sprintf("%d", t.Port)))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error listening:", err.Error())
				return
			}

			t.Programm.Send(models.ServerMsg{Conn: conn, Service: models.ServiceMsg{Type: "PrepareUser"}})

			f := func(conn *net.Conn, msg models.ServerMsg) {
				if msg.Service.Type == "HandServer" {
					t.SendMsg(conn, *t.Name, "HandServer")
				}

				t.Programm.Send(msg)

				t.Programm.Send(models.TimerCount{Data: *t.Name})
			}
			go t.handleRequest(&conn, f)
		}
	}()
}

func (t *Tcp) handleRequest(conn *net.Conn, f func(conn *net.Conn, msg models.ServerMsg)) {
	scanner := bufio.NewScanner(*conn)
	msg := models.ServerMsg{Conn: *conn}
	for scanner.Scan() {
		var s models.ServiceMsg
		err := json.Unmarshal(scanner.Bytes(), &s)
		if err != nil {
			fmt.Println("error unmarshaling message")
			return
		}
		msg.Service = s
		if msg.Service.Type == "" {
			msg.Service.Type = "Message"
		}
		f(conn, msg)
	}

}
