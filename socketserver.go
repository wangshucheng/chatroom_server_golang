package main

import (
	network "chatroom_server_golang/network"
	"fmt"
	"log"
	"net"
)

var(
	connects []net.Conn
)

func Listening(addr string) {
	tcpListen, err := net.Listen("tcp", addr)

	if err != nil {
		panic(err)
	}

	log.Println("running")

	for {
		conn, err := tcpListen.Accept()
		if err != nil {
			log.Println("accept failed, err is",err)
			continue
		}
		go connHandle(conn)

	}
}

func connHandle(conn net.Conn) {
	connects = append(connects, conn)

	defer conn.Close()
	defer onClose(conn)

	_, err := conn.Write([]byte("welcome"))
	if err != nil {
		log.Println("write welcome failed, err is",err)
		return
	}

	readBuff := make([]byte, 1024)

	for {
		n, err := conn.Read(readBuff)
		if err != nil {
			log.Println("read failed, err is",err)
			return
		}
		onRead(readBuff[:n], conn)
	}
}

func onClose(conn net.Conn) {
	log.Println("onclose",conn)
	for index, c := range connects {
		if conn == c {
			connects = append(connects[:index], connects[index+1:]...)
			break
		}
	}
}

func onRead(readBuff []byte,connect net.Conn) {
	log.Println("receive buf:",string(readBuff))

	context, err := network.ResolveMessage(readBuff)
	if err != nil {
		log.Println("resolve msg failed, err is",err)
		return
	}

	sender:=connect.RemoteAddr().String()

	fmt.Println("receive sender:", sender,", msg:",context)

	buf, err := network.MakeMessage(sender, context)
	if err != nil {
		log.Println("make msg failed, err is",err)
		return
	}

	for _, conn := range connects {
		_, err = conn.Write(buf)
		if err != nil {
			log.Println("write failed, err is",err)
		}
	}
}