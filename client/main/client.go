package main

import (
	"chatroom/client/clientprocess"
	"chatroom/comm"
	"fmt"
	"net"
	"os"
)

var loop bool = true
var key int
var username, password string
var conn net.Conn

func main() {

	// 连接服务器
	var err error
	conn, err = net.Dial("tcp", "127.0.0.1:25535")
	if err != nil {
		fmt.Printf("connect server. net.Dial error %v\n", err)
		return
	}

	for loop {
		fmt.Println("------------------------欢迎进入多人聊天系统------------------------")
		fmt.Println(" 				1、登录聊天室")
		fmt.Println("				2、注册用户")
		fmt.Println("				3、退出")
		fmt.Println("--------------------------请选择(1-3):-----------------------------")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			err = SelectUserLogin()
			if err != nil {
				fmt.Printf("err: %v\n", err)
			} else {
				ShowLoginedUI()
			}
		case 2:
			SelectUserRegister()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("请输入正确的选项")
		}
	}
}

func SelectUserLogin() (err error) {
	fmt.Println("**登录**")
	fmt.Println("请输入用户名")
	fmt.Scanf("%s\n", &username)
	fmt.Println("请输入密码")
	fmt.Scanf("%s\n", &password)

	um := clientprocess.UserManager{
		Socket: conn,
	}
	err = um.ProcessLogin(username, password)
	if err != nil {
		return
	}
	fmt.Println("**登录成功**")
	return
}

func SelectUserRegister() (err error) {
	fmt.Println("**注册**")
	fmt.Println("请输入用户名")
	fmt.Scanf("%s\n", &username)
	fmt.Println("请输入密码")
	fmt.Scanf("%s\n", &password)

	um := clientprocess.UserManager{
		Socket: conn,
	}
	err = um.UserRegister(username, password)
	if err != nil {
		fmt.Println("**注册失败**")
		return
	}
	fmt.Println("**注册成功**")
	return nil
}

func SelectUserList() error {
	fmt.Println("**在线列表**")
	um := clientprocess.UserManager{
		Socket: conn,
	}
	s, err := um.ListUser()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	for _, v := range s {
		fmt.Println(v)
	}

	return nil
}

func SelectUserPrivateChat() (err error) {
	fmt.Println("**私聊**")
	var strName string
	var strData string
	fmt.Println("请输入要私聊的用户名称：")
	fmt.Scanf("%s\n", &strName)
	fmt.Println("请输入要发送的信息：")
	fmt.Scanf("%s\n", &strData)

	// Msg->SmsMsg->SmsPrivateChat
	var SmsPriMsg comm.SmsPrivateChat
	SmsPriMsg.ChatData = strData
	SmsPriMsg.SendUserName = username
	SmsPriMsg.RecvUserName = strName

	sm := clientprocess.SmsManager{
		Socket: conn,
	}
	return sm.Sms_PrivateChat(SmsPriMsg)
}
