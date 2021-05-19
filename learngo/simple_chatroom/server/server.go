package main

import (
	"fmt"
	"net"
)

var connMap = make(map[string]*net.TCPConn)

func processMessage(conn *net.TCPConn) {
	buf := make([]byte, 1024)
	for {
		offset, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		if offset > 0 {
			fmt.Printf("%s: %s\n", conn.RemoteAddr().String(), string(buf[:offset]))
		}

		//broadcast
		for _, tcpConn := range connMap {
			if tcpConn.RemoteAddr().String() == conn.RemoteAddr().String() {
				continue
			}
			msg := string(buf[:offset])
			msg = conn.RemoteAddr().String() + ":" + msg
			_, _ = tcpConn.Write([]byte(msg))
		}
	}
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	socket, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer func(socket net.Listener) {
		err := socket.Close()
		if err != nil {
			panic(err)
		}
	}(socket)

	for {
		conn, err := socket.AcceptTCP()
		if err != nil {
			panic(err)
		}
		connMap[conn.RemoteAddr().String()] = conn
		fmt.Println("连接的客户端信息:", conn.RemoteAddr().String())
		go processMessage(conn)
	}
}
