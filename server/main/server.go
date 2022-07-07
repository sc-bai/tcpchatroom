package main

import (
	rediscache "chatroom/server/redis"
	"fmt"
	"net"
)

func init() {
	var err error
	err = rediscache.InitDb()
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
	if rediscache.DBclient != nil {
		defer rediscache.DBclient.Close()
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
	defer fmt.Println("[server]: client outline:", conn.RemoteAddr().String())
	p := Processor{
		Socket: conn,
	}
	return p.ProcessHandle()
}
