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
	f := func(conn *net.Conn, msg models.UserMessage) {
		asd := models.Message{Conn: *conn, Data: msg.Message}
		if msg.IsHandShake {
			t.Programm.Send(models.HandShake{Conn: *conn, Data: msg.Message})
			return
		}
		if msg.IsCall {
			t.Programm.Send(models.CallAction{Conn: *conn, CallStatus: msg.CallStatus})
			return
		}
		t.Programm.Send(asd)
	}
	go t.handleRequest(&conn, f)
	return &conn, err
}

func (t *Tcp) SendMsg(c *net.Conn, msg models.UserMessage) {
	conn := *c
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
			t.Programm.Send(models.PrepareUser{Conn: conn})
			f := func(conn *net.Conn, msg models.UserMessage) {
				if msg.IsHandShake {
					asd := models.UserMessage{Message: *t.Name, IsHandShake: true}
					t.SendMsg(conn, asd)
					t.Programm.Send(models.HandShake{Conn: *conn, Data: msg.Message})
					return
				}
				if msg.IsCall {
					t.Programm.Send(models.CallAction{Conn: *conn, CallStatus: msg.CallStatus})
					return
				}
				t.Programm.Send(models.Message{Conn: *conn, Data: msg.Message})
			}
			go t.handleRequest(&conn, f)
		}
	}()
}

func (t *Tcp) handleRequest(conn *net.Conn, f func(conn *net.Conn, msg models.UserMessage)) {
	scanner := bufio.NewScanner(*conn)
	for scanner.Scan() {
		var s models.UserMessage
		err := json.Unmarshal(scanner.Bytes(), &s)
		if err != nil {
			fmt.Println("error unmarshaling message")
			return
		}
		f(conn, s)
	}
}
