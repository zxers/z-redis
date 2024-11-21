package main

import (
	"log"
	"net"
	"sync"
)

type Handler interface {
	Handle(conn net.Conn)
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
