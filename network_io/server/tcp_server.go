package main

import (
	"fmt"
	"log"
	"net"
)

type TcpListener struct {
	Addr string
	Listener *net.TCPListener
}

func (l *TcpListener) Start() {
	fmt.Printf("Start listening on %s\n", l.Addr)
	for {
		// get connection
		conn, err := l.Listener.Accept()
		if err != nil {
			fmt.Println("cannot accept:", err)
			continue
		}

		// write
		go l.handleConnection(conn)
	}
}

func (l *TcpListener) handleConnection(conn net.Conn) {
	defer func() {
		log.Printf("closing current connection...")
		err := conn.Close()
		if err != nil {
			fmt.Println("connection closing failed: ", err)
		}
	}()

	str := fmt.Sprintf("Hello, net pkg: password is %s", generateRandomString(PasswordLength))
	data := []byte(str)
	n, err := conn.Write(data)
	if err != nil {
		fmt.Println("cannot write:", err)
		return
	}

	log.Printf("write %d bytes\n", n)
}

func NewListener(addr string) *TcpListener {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	tcpListener, ok := ln.(*net.TCPListener)
	if !ok {
		panic("cannot convert to TCPListener")
	}

	return &TcpListener{
		Addr: addr,
		Listener: tcpListener,
	}
}

func main() {
	addr := "localhost:8081"
	l := NewListener(addr)
	l.Start()
	defer func() {
		log.Println("closing listener...")
		err := l.Listener.Close()
		if err != nil {
			fmt.Println("connection closing failed: ", err)
		}
	}()
}