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
			//fmt.Printf("[server]: readpage error: %v\n", err)
			return err
		}

		switch msg.Code {
		case comm.CodeLogin:
			fmt.Println("[server]: ccmd user login.")
			u := process.UserProcess{
				Socket: p.Socket,
			}
			err = u.UserLogin(&msg)
		case comm.CodeRegister:
			fmt.Println("[server]: cmd user register.")
			u := process.UserProcess{
				Socket: p.Socket,
			}
			err2 := u.UserRegister(&msg)
			if err2 != nil {
				return err2
			}
		case comm.CodeSms:
			fmt.Println("sms")
		}
	}
}
