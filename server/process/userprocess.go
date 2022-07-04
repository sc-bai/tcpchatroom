package process

import (
	"chatroom/comm"
	rediscache "chatroom/server/redis"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Socket  net.Conn
	Redisdb rediscache.RedisDb
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

// 还有问题
func (p *UserProcess) UserRegister(msg *comm.Msg) (err error) {

	var registmsg comm.RegistMsg
	err = json.Unmarshal([]byte(msg.Data), &registmsg)
	if err != nil {
		fmt.Printf("err2: %v\n", err)
		return
	}

	fmt.Printf("[server]: registmsg.UserName: %v registmsg.UserPasswd: %v\n", registmsg.UserName, registmsg.UserPasswd)

	lmg, err2 := p.Redisdb.FindUser(p.Socket.RemoteAddr().String(), registmsg.UserName)
	if err2 != nil {
		return err2
	}

	var regRes comm.ServerRes
	if len(lmg.UserId) > 0 {
		// 说明存在
		regRes.Code = comm.ServerFail
		regRes.Msg = "this name has alread registered"
	} else {
		// 注册成功 写入缓存
		p.Redisdb.PutUser(comm.LoginMessage{
			UserId:     registmsg.UserName + "_id",
			UserName:   registmsg.UserName,
			UserPasswd: registmsg.UserPasswd,
		}, p.Socket.RemoteAddr().String())

		regRes.Code = comm.ServerSuccess
		regRes.Msg = "register success"
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
