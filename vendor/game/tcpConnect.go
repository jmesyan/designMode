package main

import (
	"net"
)

type tcpConnect struct {
	session *tcpSession
}

func newTCPConnect(addr string, readChan chan tcpReadMessage) *tcpConnect {
	tcpConnect := new(tcpConnect)
	tcpConnect.start(addr, readChan)
	return tcpConnect
}

func (t *tcpConnect) start(addr string, readChan chan tcpReadMessage) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if nil != err {
		return
	}
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if nil != err {
		return
	}
	t.session = newTCPSession(conn, readChan)
}

func (t *tcpConnect) close() {
	t.session.close()
}
