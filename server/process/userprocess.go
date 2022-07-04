package process

import (
	"chatroom/comm"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Socket net.Conn
}

func (p *UserProcess) UserLogin(msg *comm.Msg) (err error) {

	var loginMessage comm.LoginMessage
	err = json.Unmarshal([]byte(msg.Data), &loginMessage)
	if err != nil {
		fmt.Printf("err2: %v\n", err)
		return
	}

	fmt.Printf("loginMessage.UserId: %v\n", loginMessage.UserId)
	fmt.Printf("loginMessage.UserName: %v\n", loginMessage.UserName)
	fmt.Printf("loginMessage.UserPasswd: %v\n", loginMessage.UserPasswd)

	return nil
}

func (p *UserProcess) UserRegister(msg *comm.Msg) (err error) {

	return nil
}
