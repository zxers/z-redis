package main

import "strconv"

const CRLF = "\r\n"

type Message interface {
	ToBytes() []byte
}

type SimpleString struct {
	Str string
}

func (s *SimpleString) ToBytes() []byte {
	return []byte("+" + s.Str + CRLF)
}

func NewSimpleString(str string) *SimpleString {
	return &SimpleString{
		Str: str,
	}
}

type ErrMsg struct {
	Err string
}

func (e *ErrMsg) ToBytes() []byte {
	return []byte("-" + e.Err + CRLF)
}

func NewErrMsg(err string) *ErrMsg {
	return &ErrMsg{
		Err: err,
	}
}

type Array struct {
	Arg [][]byte
}

func (a *Array) ToBytes() []byte {
	res := "*"
	argLen := len(a.Arg)
	argLenStr := strconv.Itoa(argLen)
	res += argLenStr + CRLF
	for _, val := range a.Arg {
		argLen = len(val)
		argLenStr := strconv.Itoa(argLen)
		res += "$" + argLenStr + CRLF + string(val) + CRLF
	}
	return []byte(res)
}

func NewArray(arg [][]byte) *Array {
	return &Array{
		Arg: arg,
	}
}
