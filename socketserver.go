package main

import (
	"fmt"
	"log"
	"net"
)

func Listening() {
	tcpListen, err := net.Listen("tcp", ":1238")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := tcpListen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go connHandle(conn)
	}
}
func  connHandle(conn net.Conn) {
	defer conn.Close()
	readBuff := make([]byte, 1024)
	for {
		n, err := conn.Read(readBuff)
		if err != nil {
			log.Println(err)
			return
		}

		conn.Write([]byte("hello too\n"))
		fmt.Println(string(readBuff[:n]))

	}
}