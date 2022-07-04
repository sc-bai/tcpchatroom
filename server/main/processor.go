package main

import (
	"chatroom/comm"
	"chatroom/server/process"
	"fmt"
	"net"
)

type Processor struct {
	Socket net.Conn
	Data   []byte
}

func (p *Processor) ProcessHandle() (err error) {
	for {
		t := &comm.Transfer{
			Sock: p.Socket,
		}
		msg, err := t.ReadPkg()
		if err != nil {
			fmt.Printf("[server]: readpage error: %v\n", err)
			return err
		}

		switch msg.Code {
		case comm.CodeLogin:
			fmt.Println("[server]:user login.")
			u := process.UserProcess{
				Socket: p.Socket,
			}
			u.UserLogin(&msg)
		case comm.CodeRegister:
			fmt.Println("[server]:user register.")
			u := process.UserProcess{
				Socket: p.Socket,
			}
			u.UserRegister(&msg)
		case comm.CodeSms:
			fmt.Println("sms")
		}
	}
}
