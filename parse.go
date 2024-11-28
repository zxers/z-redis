package main

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

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
