package main

import (
	"net"
	"strconv"
)

const (
	readMessageFlag = iota
	writeMessageFlag
	stopFlag
)

type tcpServerMessage struct {
	flag    int
	context []byte
}

type tcpReadMessage struct {
	remoteIP string
	localIP  string
	header   MessageHeader
	body     []byte
}

type tcpServer struct {
	port     int
	sessions sessionPool
	readChan chan tcpReadMessage
}

func newTCPServer(port int, readChan chan tcpReadMessage) *tcpServer {
	newServer := new(tcpServer)
	newServer.port = port
	newServer.readChan = readChan
	newServer.start()
	return newServer
}

func (t *tcpServer) start() {
	portStr := ":" + strconv.Itoa(t.port)
	tcp4Addr, err := net.ResolveTCPAddr("tcp4", portStr)
	if nil != err {
		spruceLog.error("%v", err)
		return
	}
	listen, err := net.ListenTCP("tcp", tcp4Addr)
	if nil != err {
		spruceLog.error("%v", err)
		return
	}
	t.sessions.init()
	go t.startAccpet(listen)
}

func (t *tcpServer) startAccpet(listen *net.TCPListener) {
	if nil == listen {
		return
	}
	spruceLog.info("start listen tcp4 at port :%d ...", t.port)
	for {
		conn, err := listen.AcceptTCP()
		if nil != err {
			spruceLog.error("accpet error:%v", conn.RemoteAddr().String(), err)
			continue
		}
		t.sessions.addNewTCPConnect(conn, t.readChan)
	}
}
