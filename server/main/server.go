package main

import (
	rediscache "chatroom/server/redis"
	"fmt"
	"net"

	"github.com/go-redis/redis"
)

var g_db *redis.Client

func init() {
	var err error
	g_db, err = rediscache.InitDb()
	if err != nil {
		fmt.Println("[server]initdb error", err)
	}
}
func main() {

	l, err := net.Listen("tcp", "0.0.0.0:25535")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer l.Close()
	if g_db != nil {
		defer g_db.Close()
	}

	fmt.Println("[server]: start server success...")

	for {
		fmt.Println("[server]: wait for connect...")
		c, err := l.Accept()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}
		fmt.Printf("[server]: client connected, ip: %v\n", c.RemoteAddr().String())
		go HandleConnect(c)
	}
}

func HandleConnect(conn net.Conn) (err error) {
	defer conn.Close()
	p := Processor{
		Socket:  conn,
		redisDb: g_db,
	}
	return p.ProcessHandle()
}
