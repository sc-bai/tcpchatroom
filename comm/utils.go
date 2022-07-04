package comm

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

/*
	这里是从sock中读写数据
*/
type Transfer struct {
	Sock net.Conn
	Data []byte
}

func (p *Transfer) ReadPkg() (msg Msg, err error) {
	p.Data = make([]byte, 1024)
	n, err := p.Sock.Read(p.Data[:4]) // 先发送四个字节存档长度
	if err != nil {
		fmt.Printf("Socket.Read error: %v\n", err)
		return
	}

	u := binary.BigEndian.Uint32(p.Data[:n]) // 四个字节长度切片转uint32

	n2, err2 := p.Sock.Read(p.Data[:u])
	if uint32(n2) != u || err2 != nil {
		fmt.Printf("Sock.Read 2: %v\n", err2)
		err = err2
		return
	}

	err3 := json.Unmarshal(p.Data[:n2], &msg)
	if err3 != nil {
		fmt.Printf("json.Unmarshal: %v\n", err3)
		err = err3
		return
	}

	return msg, nil
}

func (p *Transfer) WritePkg(msg Msg) error {

	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("WritePkg json.Marshal: %v\n", err)
		return err
	}
	u := uint32(len(b))
	fmt.Printf("WritePkg len: %v\n", u)

	// 数字转为切片发送长度
	head := make([]byte, 4)
	binary.BigEndian.PutUint32(head, u)

	n, err2 := p.Sock.Write(head)
	if n != 4 || err2 != nil {
		fmt.Printf("Sock.Write head error: %v len:%d \n", err2, n)
		return err2
	}

	n, err2 = p.Sock.Write(b)
	if err2 != nil {
		fmt.Printf("Sock.Write len error: %v\n", err2)
		return err2
	}

	fmt.Printf("client send len: %v\n", n)
	return nil
}

func (p *Transfer) WritePkgEx(b []byte) error {
	u := uint32(len(b))
	fmt.Printf("WritePkgEx len: %v\n", u)

	// 数字转为切片发送长度
	head := make([]byte, 4)
	binary.BigEndian.PutUint32(head, u)

	n, err2 := p.Sock.Write(head)
	if n != 4 || err2 != nil {
		fmt.Printf("WritePkgEx Sock.Write head error: %v len:%d \n", err2, n)
		return err2
	}

	n, err2 = p.Sock.Write(b)
	if err2 != nil {
		fmt.Printf("WritePkgEx Sock.Write len error: %v\n", err2)
		return err2
	}

	fmt.Printf("client send len: %v\n", n)
	return nil
}
