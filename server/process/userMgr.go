package process

import (
	"chatroom/comm"
	"errors"
	"fmt"
	"net"
)

// 保存在线人员信息
var UserManager *UserMgr

type UserMgr struct {
	onlineUsers map[string]*UserProcess // string 为username
}

func init() {
	UserManager = &UserMgr{
		onlineUsers: make(map[string]*UserProcess, comm.MAXONLINEUSER),
	}
}

func (p *UserMgr) AddOnlineUser(u *UserProcess) {
	fmt.Println("[server]addonlineuser", u.UserName)
	p.onlineUsers[u.UserName] = u
}

func (p *UserMgr) DeleteOnlineUser(u *UserProcess) {
	delete(p.onlineUsers, u.UserName)
}

func (p *UserMgr) DeleteOnlineUserWithConn(c net.Conn) {
	for k, v := range p.onlineUsers {
		if v.Socket == c {
			delete(p.onlineUsers, k)
			break
		}
	}
}
func (p *UserMgr) FindUserOnline(UserName string) (u *UserProcess, err error) {
	for k, v := range p.onlineUsers {
		if k == UserName {
			return v, nil
		}
	}
	return nil, errors.New("[server]online not find")
}
