package main

import (
	"fmt"
	"net"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:25535")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
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
		Socket: conn,
	}
	return p.ProcessHandle()
}
