package main

import (
	"chatroom/comm"
	"chatroom/server/process"
	rediscache "chatroom/server/redis"
	"fmt"
	"net"

	"github.com/go-redis/redis"
)

type Processor struct {
	Socket  net.Conn
	Data    []byte
	redisDb *redis.Client
}

func (p *Processor) ProcessHandle() (err error) {
	for {
		t := &comm.Transfer{
			Sock: p.Socket,
		}
		msg, err := t.ReadPkg()
		if err != nil {
			fmt.Printf("[server]: readpage error: %v\n", err)
			return err
		}

		var rdb rediscache.RedisDb
		rdb.DBclient = g_db

		switch msg.Code {
		case comm.CodeLogin:
			fmt.Println("[server]:user login.")
			u := process.UserProcess{
				Socket:  p.Socket,
				Redisdb: rdb,
			}
			u.UserLogin(&msg)
		case comm.CodeRegister:
			fmt.Println("[server]:user register.")
			u := process.UserProcess{
				Socket:  p.Socket,
				Redisdb: rdb,
			}
			u.UserRegister(&msg)
		case comm.CodeSms:
			fmt.Println("sms")
		}
	}
}
