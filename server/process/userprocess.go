package process

import (
	"chatroom/comm"
	rediscache "chatroom/server/redis"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type UserProcess struct {
	Socket   net.Conn
	UserName string
}

func (p *UserProcess) UserLogin(msg *comm.Msg) (err error) {

	var loginMessage comm.LoginMessage
	err = json.Unmarshal([]byte(msg.Data), &loginMessage)
	if err != nil {
		fmt.Printf("err2: %v\n", err)
		return
	}

	fmt.Printf("loginMessage.UserId: %v loginMessage.UserName: %v loginMessage.UserPasswd: %v\n", loginMessage.UserId, loginMessage.UserName, loginMessage.UserPasswd)

	// 查询服务器是否有缓存并且一致 就可以返回登录成功
	stRedis := rediscache.RedisDb{}
	lm := stRedis.FindUser(loginMessage.UserName)

	var regRes comm.ServerRes
	if len(lm.UserId) > 0 && lm.UserPasswd == loginMessage.UserPasswd {
		fmt.Println("login success")
		regRes.Code = comm.ServerSuccess

		// 塞进监管账号
		p.UserName = loginMessage.UserName

	} else if lm.UserPasswd != loginMessage.UserPasswd {
		fmt.Println("login error1")
		fmt.Printf("lm.UserPasswd: %v end\n", lm.UserPasswd)
		fmt.Printf("loginMessage.UserPasswd: %v end\n", loginMessage.UserPasswd)

		regRes.Code = comm.ServerFail
		regRes.Msg = "passwd error.try again"
	} else {
		regRes.Code = comm.ServerFail
		regRes.Msg = "not reigster, please register first"
	}

	var sendMsg comm.Msg
	sendMsg.Code = comm.CodeLoginRes
	b, _ := json.Marshal(&regRes)
	sendMsg.Data = string(b)

	t := comm.Transfer{
		Sock: p.Socket,
	}
	err2 := t.WritePkg(sendMsg)
	if err2 != nil {
		return err2
	}

	if regRes.Code == comm.ServerSuccess {

		UserManager.AddOnlineUser(p)

		// 通知上线
		sms := SmsMessage{}
		err3 := sms.NotifyUserOnline(loginMessage.UserName)
		if err3 != nil {
			fmt.Printf("NotifyUserOnline error: %v\n", err3)
		}
	}

	return nil
}

func (p *UserProcess) UserRegister(msg *comm.Msg) (err error) {

	var registmsg comm.RegistMsg
	err = json.Unmarshal([]byte(msg.Data), &registmsg)
	if err != nil {
		fmt.Printf("err2: %v\n", err)
		return
	}

	fmt.Printf("[server]: registmsg.UserName: %v registmsg.UserPasswd: %v\n", registmsg.UserName, registmsg.UserPasswd)

	stRedis := rediscache.RedisDb{}
	lmg := stRedis.FindUser(registmsg.UserName)

	var regRes comm.ServerRes
	if len(lmg.UserId) > 0 {
		// 说明存在
		regRes.Code = comm.ServerFail
		regRes.Msg = "this name has alread registered"
		fmt.Println("[server] register failed")
	} else {
		// 注册成功 写入缓存
		stRedis.PutUser(comm.LoginMessage{
			UserId:     registmsg.UserName + "_id",
			UserName:   registmsg.UserName,
			UserPasswd: registmsg.UserPasswd,
		})

		regRes.Code = comm.ServerSuccess
		regRes.Msg = "register success"
		fmt.Println("[server] register success")
	}

	b, err3 := json.Marshal(regRes)
	if err3 != nil {
		fmt.Printf("UserRegister Marshal err: %v\n", err3)
		return err3
	}

	msgRet := comm.Msg{
		Code: comm.CodeRegisterRes,
		Data: string(b),
	}

	trans := comm.Transfer{
		Sock: p.Socket,
	}
	err = trans.WritePkg(msgRet)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	return nil
}

func (p *UserProcess) UserList(msg *comm.Msg) error {
	if msg.Code != comm.CodeUserList {
		return errors.New("code type error")
	}

	var data string
	for _, v := range UserManager.onlineUsers {
		data += v.UserName
		data += "-"
	}

	if len(data) > 0 {
		data = data[0:(len(data) - 1)]
	}

	msgRet := comm.Msg{
		Code: comm.CodeUserListRes,
		Data: data,
	}
	trans := comm.Transfer{
		Sock: p.Socket,
	}
	return trans.WritePkg(msgRet)
}
