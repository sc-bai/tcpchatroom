package main

import (
	"chatroom/client/clientprocess"
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
		fmt.Println("------------欢迎进入多人聊天系统------------")
		fmt.Println("------------1、登录聊天室")
		fmt.Println("------------2、注册用户")
		fmt.Println("------------3、退出")
		fmt.Println("------------请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			SelectUserLogin()
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

	return nil
}
