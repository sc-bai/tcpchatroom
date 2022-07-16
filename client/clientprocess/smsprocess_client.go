package clientprocess

import (
	"chatroom/comm"
	"encoding/json"
	"net"
)

type SmsManager struct {
	Socket net.Conn
}

func (p *SmsManager) Sms_PrivateChat(stchat comm.SmsPrivateChat) error {
	// // Msg->SmsMsg->SmsPrivateChat

	var msg comm.Msg
	var smsMsg comm.SmsMsg
	smsMsg.Code = comm.CodeSmsPrivateChat
	b, _ := json.Marshal(stchat)
	smsMsg.SmsData = string(b)
	b, _ = json.Marshal(smsMsg)
	msg.Code = comm.CodeSms
	msg.Data = string(b)

	t := comm.Transfer{
		Sock: p.Socket,
	}

	// 这里应该做个接受发送私聊消息的通知
	return t.WritePkg(msg)
}
