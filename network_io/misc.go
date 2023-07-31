package main

import (
	"net"
	"time"
)


type Conn interface {
    Read(b []byte) (n int, err error)
    Write(b []byte) (n int, err error)
    Close() error
    LocalAddr() net.Addr
    RemoteAddr() net.Addr
    SetDeadline(t time.Time) error
    SetReadDeadline(t time.Time) error
    SetWriteDeadline(t time.Time) error
}

var _ net.Conn = (*net.TCPConn)(nil)
var _ net.Conn = (*net.UDPConn)(nil)
var _ net.Conn = (*net.IPConn)(nil)
var _ net.Conn = (*net.UnixConn)(nil)

func testNetworkIo() {
	// fmt.Printf("net.Conn: %T\n", net.Conn)
}