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

			// 客户端断开连接
			process.UserManager.DeleteOnlineUserWithConn(p.Socket)

			return err
		}

		switch msg.Code {
		case comm.CodeLogin:
			fmt.Println("[server]: ccmd user login.")
			u := process.UserProcess{
				Socket: p.Socket,
			}
			err = u.UserLogin(&msg)
			if err != nil {
				fmt.Printf("[server] Userlogin error: %v\n", err)
			}
		case comm.CodeRegister:
			fmt.Println("[server]: cmd user register.")
			u := process.UserProcess{
				Socket: p.Socket,
			}
			err2 := u.UserRegister(&msg)
			if err2 != nil {
				fmt.Printf("[server]: UserRegister error. %v\n", err2)
			}

		case comm.CodeUserList:
			fmt.Println("[server]: cmd user lsit.")

			u := process.UserProcess{
				Socket: p.Socket,
			}

			err = u.UserList(&msg)
			if err != nil {
				fmt.Printf("[server]: userlist error,%v\n", err)
			}

		case comm.CodeSms:
			s := process.SmsMessage{}
			err = s.HandleSmsMessage(msg)
			if err != nil {
				fmt.Println("[server]:handle sns message error")
			}
		}
	}
}
