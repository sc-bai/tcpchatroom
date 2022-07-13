package main

import "fmt"

var secondkey int

//  二级菜单---登录之后的界面
func ShowLoginedUI() {
	for {
		fmt.Println("		^^^^^^^^^^^^^^^^^登录成功^^^^^^^^^^^^^^")
		fmt.Println("\t\t\t\t1、查看在线人员列表")
		fmt.Println("\t\t\t\t2、请输入要私聊的名称")
		fmt.Println("\t\t\t\t3、广播")
		fmt.Println("\t\t\t\t4、退出")

		fmt.Scanf("%d\n", &secondkey)
		switch secondkey {
		case 1:
			fmt.Println("查看在线人员列表")
			SelectUserList()
		case 2:
			fmt.Println("请输入要私聊的名称")

		case 3:
			fmt.Println("广播")
		case 4:
			fmt.Println("退出二级菜单")
			return
		default:
			fmt.Println("请输入有效的数字")
		}
	}
}
