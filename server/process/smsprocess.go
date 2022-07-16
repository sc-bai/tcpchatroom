package process

import (
	"chatroom/comm"
	"encoding/json"
	"errors"
	"fmt"
)

type SmsMessage struct {
}

func (P *SmsMessage) HandleSmsMessage(msg comm.Msg) (err error) {

	if msg.Code != comm.CodeSms {
		return errors.New("sms tpye error")
	}
	var smsMsg comm.SmsMsg
	json.Unmarshal([]byte(msg.Data), &smsMsg)

	if smsMsg.Code == comm.CodeSmsPrivateChat {
		err = P.ChatPrivate(smsMsg)
		return err
	} else if smsMsg.Code == comm.CodeSmsGorup {
		err = P.NotifyAllOnlineUser(smsMsg)
		return err
	}

	return nil
}

// 通知别人自己上线
func (p *SmsMessage) NotifyUserOnline(username string) error {
	fmt.Println("notify user online")
	smsdata := comm.SmsMsg{
		Code:    comm.CodeSmsGroupRes,
		SmsData: "[group]" + username + " online now!!",
	}
	return p.NotifyAllOnlineUser(smsdata)
}

// 群聊 data 发送数据
func (p *SmsMessage) NotifyAllOnlineUser(smsdata comm.SmsMsg) error {
	var msg comm.Msg
	msg.Code = comm.CodeSms

	b, _ := json.Marshal(smsdata)
	msg.Data = string(b)

	for _, v := range UserManager.onlineUsers {

		t := comm.Transfer{
			Sock: v.Socket,
		}

		err := t.WritePkg(msg)
		if err != nil {
			fmt.Printf("NotifyAllOnlineUser error: %v\n", err)
		}
	}

	return nil
}

func (p *SmsMessage) ChatPrivate(smsMsg comm.SmsMsg) error {

	// Msg->SmsMsg->SmsPrivateChat

	var prvMsg comm.SmsPrivateChat
	json.Unmarshal([]byte(smsMsg.SmsData), &prvMsg)

	var msgRes comm.Msg
	var msgSmsR comm.SmsMsg
	msgRes.Code = comm.CodeSms
	// 解析出来要发送的对象和数据
	up, err := UserManager.FindUserOnline(prvMsg.SendUserName)
	if err != nil {
		// 返回失败
		fmt.Printf(":FindUserOnlineByName err %v\n", err)
		// 不做错误处理了  应该返回发送失败
		return err
	}

	t := comm.Transfer{
		Sock: up.Socket,
	}

	msgSmsR.Code = comm.CodeSmsPrivateChat
	msgSmsR.SmsData = smsMsg.SmsData // 发过来的数据 直接给另外一个客户端

	b, _ := json.Marshal(msgSmsR)
	msgRes.Data = string(b)

	t.WritePkg(msgRes)

	return nil
}
