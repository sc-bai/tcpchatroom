package clientprocess

import (
	"chatroom/comm"
	"encoding/json"
	"fmt"
	"net"
)

// var g_conn net.Conn

type UserManager struct {
	Socket net.Conn
}

func PrintProcess() {
	fmt.Println("printfprocess")
}
func (p *UserManager) ProcessLogin(name, passwd string) error {

	// 发送到服务器
	loginMsg := comm.LoginMessage{
		UserName:   name,
		UserPasswd: passwd,
	}

	b, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Printf("ProcessLogin:json.Marshal: %v\n", err)
		return err
	}

	msg := comm.Msg{
		Code: comm.CodeLogin,
		Data: string(b),
	}

	t := comm.Transfer{
		Sock: p.Socket,
	}
	err = t.WritePkg(msg)
	if err != nil {
		fmt.Printf("ProcessLogin WritePkg: %v\n", err)
		return err
	}

	// 接收服务器返回
	// todo ...
	return nil
}
