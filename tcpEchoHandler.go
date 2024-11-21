package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

type EchoHandler struct {
}

func (e *EchoHandler) Handle(conn net.Conn) {
	rd := bufio.NewReader(conn)
	for {
		readString, err := rd.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				break
			}
		}
		conn.Write([]byte("echo: " + readString))
	}
	err := conn.Close()
	if err != nil {
		log.Println(err)
	}
}
