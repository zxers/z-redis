package main

import (
	"errors"
	"sync"
)

type Command [][]byte

var DBInstance *DB

type DB struct {
	dict Dict
}

func NewDB() *DB {
	return &DB{
		dict: NewSyncMapDict(),
	}
}

// Exec arg exp: [][]byte{[]byte("set"), []byte("key"), []byte("value")}
func (db *DB) Exec(arg Command) Message {
	cmd := string(arg[0])
	exec, ok := commandTable[cmd]
	if !ok {
		return NewErrMsg("command not found")
	}
	return exec(db, arg[1:])
}

type Dict interface {
	Get(key string) (interface{}, error)
	Put(key string, value interface{}) error
}

type SyncMapDict struct {
	Data sync.Map
}

func NewSyncMapDict() *SyncMapDict {
	return &SyncMapDict{
		Data: sync.Map{},
	}
}

func (s *SyncMapDict) Get(key string) (interface{}, error) {
	value, ok := s.Data.Load(key)
	if !ok {
		return nil, errors.New("key is not exist")
	}
	return value, nil
}

func (s *SyncMapDict) Put(key string, value interface{}) error {
	s.Data.Store(key, value)
	return nil
}
