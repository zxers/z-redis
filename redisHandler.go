package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
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
		fmt.Println(string(array.Arg[0]), string(array.Arg[1]), string(array.Arg[2]))
		reply := DBInstance.Exec(array.Arg)
		conn.Write(reply.ToBytes())
	}
	err := conn.Close()
	if err != nil {
		log.Println(err)
	}
}

type Command [][]byte

func Parse(reader *bufio.Reader) (Message, error) {
	head, err := reader.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			return nil, err
		}
	}
	head = strings.TrimSuffix(head, CRLF)
	switch head[0] {
	case '*':
		message, err := parseArray(head[1:], reader)
		if err != nil {
			return nil, err
		}
		return message, nil
	default:
		return nil, errors.New("parse error")
	}
}

// parseArray *3\r\n$3\r\nset\r\n$3\r\nkey\r\n$5\r\nvalue\r\n
func parseArray(head string, reader *bufio.Reader) (Message, error) {
	bodyLen, err := strconv.Atoi(head)
	if err != nil {
		return nil, err
	}
	arg := make([][]byte, bodyLen)
	for i := 0; i < bodyLen; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimSuffix(line, CRLF)
		dateLen, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}
		buf := make([]byte, dateLen+2)
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return nil, err
		}
		arg[i] = buf[:len(buf)-2]
	}
	return NewArray(arg), nil
}

var DBInstance *DB

type DB struct {
}

func NewDB() *DB {
	return &DB{}
}

// Exec arg exp: [][]byte{[]byte("set"), []byte("key"), []byte("value")}
func (db *DB) Exec(arg Command) Message {
	return NewArray(arg)
}
