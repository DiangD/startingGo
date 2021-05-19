package main

import (
	"fmt"
	"net"
	"os"
)

var ch = make(chan interface{})

func reader(conn *net.TCPConn) {
	buf := make([]byte, 1024)
	for true {
		offset, err := conn.Read(buf)
		if err != nil {
			ch <- struct{}{}
			break
		}
		fmt.Println(string(buf[:offset]))
	}
}

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Server is not starting")
		panic(err)
	}
	defer func(conn *net.TCPConn) {
		_ = conn.Close()
	}(conn)

	go reader(conn)

	for {
		var msg string
		_, _ = fmt.Scanf("%s", &msg)
		_, _ = conn.Write([]byte(msg))

		select {
		case <-ch:
			fmt.Println("Server错误!请重新连接!")
			os.Exit(1)
		default:

		}
	}
}
