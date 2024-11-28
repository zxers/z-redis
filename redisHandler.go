package main

import (
	"bufio"
	"log"
	"net"
)

type RedisHandler struct {
}

func (r *RedisHandler) Handle(conn net.Conn) {
	rd := bufio.NewReader(conn)
	for {
		message, err := Parse(rd)
		if err != nil {
			continue
		}
		array, ok := message.(*Array)
		if !ok {
			continue
		}
		// fmt.Println(string(array.Arg[0]), string(array.Arg[1]), string(array.Arg[2]))
		reply := DBInstance.Exec(array.Arg)
		conn.Write(reply.ToBytes())
	}
	err := conn.Close()
	if err != nil {
		log.Println(err)
	}
}
