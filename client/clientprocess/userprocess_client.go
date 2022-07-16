package clientprocess

import (
	"chatroom/comm"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strings"
)

// var g_conn net.Conn

type UserManager struct {
	Socket net.Conn
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

	m, err2 := t.ReadPkg()
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
		return err2
	}

	if m.Code != comm.CodeLoginRes {
		return errors.New("type not right")
	}

	var res comm.ServerRes
	json.Unmarshal([]byte(m.Data), &res)

	if res.Code == comm.ServerSuccess {
		return nil
	} else {
		fmt.Printf("register ret: res.Msg: %v\n", res.Msg)
		return errors.New(res.Msg)
	}
}

func (p *UserManager) UserRegister(name, passwd string) error {

	var msg comm.Msg
	msg.Code = comm.CodeRegister

	reg := &comm.RegistMsg{
		UserName:   name,
		UserPasswd: passwd,
	}

	b, err := json.Marshal(reg)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	msg.Data = string(b)
	t := comm.Transfer{
		Sock: p.Socket,
	}
	err = t.WritePkg(msg)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	m, err2 := t.ReadPkg()
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
		return err2
	}

	if m.Code != comm.CodeRegisterRes {
		return errors.New("type not right")
	}

	var res comm.ServerRes
	json.Unmarshal([]byte(m.Data), &res)

	if res.Code == comm.ServerSuccess {
		return nil
	} else {
		fmt.Printf("register ret: res.Msg: %v\n", res.Msg)
		return errors.New(res.Msg)
	}
}

func (p *UserManager) ListUser() ([]string, error) {

	var ret []string
	var err error
	var msg comm.Msg
	reg := &comm.RegistMsg{
		UserName:   "test",
		UserPasswd: "passwd",
	}

	b, _ := json.Marshal(reg)
	/* msg := comm.Msg{
		Code: comm.CodeUserList,
		Data: string(b),
	} */
	msg.Code = comm.CodeUserList
	msg.Data = string(b)
	t := comm.Transfer{
		Sock: p.Socket,
	}

	err2 := t.WritePkg(msg)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
		return ret, err2
	}
	m, err := t.ReadPkg()
	if m.Code != comm.CodeUserListRes {
		return ret, err
	}
	fmt.Printf("m.Data: %v\n", m.Data)
	ret = strings.Split(m.Data, "-")
	return ret, err
}
