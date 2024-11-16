package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"sync"
)

type Handler interface {
	Handle(conn net.Conn)
}

func main() {
	ListenAndServe(":3007", &EchoHandler{})
}

func ListenAndServe(addr string, handler Handler) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	log.Println("Listening on ", addr)
	wg := sync.WaitGroup{}
	for {
		conn, err := listen.Accept()
		log.Println("获取新连接")
		if err != nil {
			break
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			handler.Handle(conn)
		}()
	}
	wg.Wait()
}

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
