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
	go t.ClienthandleRequest(conn)
	return &conn, err
}

func (t *Tcp) ClienthandleRequest(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	msg := models.ServerMsg{Conn: conn}
	for scanner.Scan() {
		var s models.ServiceMsg
		err := json.Unmarshal(scanner.Bytes(), &s)
		if err != nil {
			fmt.Println("error unmarshaling message")
			return
		}
		msg.Text = s
		t.Programm.Send(msg)
	}
}

func (t *Tcp) SendMsg(c *net.Conn, text, service string) {
	conn := *c
	msg := models.ServiceMsg{Service: service, Msg: text}
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
			t.Programm.Send(models.ServerMsg{Conn: conn, Text: models.ServiceMsg{Service: "PrepareUser"}})
			go t.ServerhandleRequest(&conn)
		}
	}()
}

func (t *Tcp) ServerhandleRequest(conn *net.Conn) {
	scanner := bufio.NewScanner(*conn)
	msg := models.ServerMsg{Conn: *conn}
	for scanner.Scan() {
		var s models.ServiceMsg
		err := json.Unmarshal(scanner.Bytes(), &s)
		if err != nil {
			fmt.Println("error unmarshaling message")
			return
		}
		msg.Text = s
		if s.Service == "HandServer" {
			t.SendMsg(conn, *t.Name, "HandServer")
		}
		t.Programm.Send(msg)
	}

}
