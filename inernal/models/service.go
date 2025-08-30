package models

import "net"

type ServiceMsg interface {
	isServiceMsg()
}

type HandShake Payload

func (s HandShake) isServiceMsg() {}

type PrepareUser Payload

func (s PrepareUser) isServiceMsg() {}

type TimerCount Payload

func (s TimerCount) isServiceMsg() {}

type ConnTimeout Payload

func (s ConnTimeout) isServiceMsg() {}

type CallAction Payload

func (s CallAction) isServiceMsg() {}

type Message Payload

func (s Message) isServiceMsg() {}

type Payload struct {
	Conn net.Conn
	Data string
	//0 - Call
	//1 - Accept
	//2 - End call or decline
	CallStatus int
}
