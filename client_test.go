package main

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:3007")
	if err != nil {
		t.Error(err)
	}
	conn.Write([]byte("*3\r\n$7\r\n\\r\\nset\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"))
	t.Log("write succeed")
	reader := bufio.NewReader(conn)
	for {
		line, _ := reader.ReadString('\n')
		fmt.Printf(line)
	}
}
