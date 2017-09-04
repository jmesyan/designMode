package main

import (
	"net"
	"sync"
)

type sessionPool struct {
	maxIndex int64
	mutex    sync.Mutex
	sessions map[int64]iConnect
}

func (s *sessionPool) init() {
	s.sessions = make(map[int64]iConnect)
}

func (s *sessionPool) addNewTCPConnect(conn *net.TCPConn, readChan chan tcpReadMessage) {
	if nil != conn {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.maxIndex++
		s.sessions[s.maxIndex] = newTCPSession(conn, readChan)
	}
}
