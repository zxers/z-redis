package main

import (
	"strings"
)

var commandTable map[string]CommandFunc

type CommandFunc func(db *DB, cmd Command) Message

func getString(db *DB, cmd Command) Message {
	key := string(cmd[0])
	val, err := db.dict.Get(key)
	if err != nil {
		return NewErrMsg(err.Error())
	}
	str, ok := val.(string)
	if !ok {
		return NewErrMsg("Data type error")
	}
	return NewSimpleString(str)
}

func setString(db *DB, cmd Command) Message {
	key := string(cmd[0])
	err := db.dict.Put(key, string(cmd[1]))
	if err != nil {
		return NewErrMsg(err.Error())
	}
	return NewSimpleString("OK")
}

func RegisterCommand(name string, cmd CommandFunc) {
	// 命令全部转为小写
	name = strings.ToLower(name)
	commandTable[name] = cmd
}

func init() {
	commandTable = make(map[string]CommandFunc)
	RegisterCommand("Get", getString)
	RegisterCommand("Set", setString)
}
