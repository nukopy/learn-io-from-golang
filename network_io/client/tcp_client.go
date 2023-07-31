package main

import (
	"log"
	"net"
	"time"
)

func connect() {
	// get connection
	addr := "localhost:8081"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("cannot connect:", err)
	}
	defer func() {
		log.Printf("closing current connection...")
		err := conn.Close()
		if err != nil {
			log.Println("connection closing failed: ", err)
		}
	}()

	// read
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("cannot read:", err)
	}
	data := string(buf[:n])
	log.Printf("read %d bytes: %s\n", n, data)
}

func connectOnce() {
	connect()
}

func connectLoop() {
	for i := 0; i < 100; i++ {
		connect()

		// wait
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// connectOnce()
	connectLoop()
}