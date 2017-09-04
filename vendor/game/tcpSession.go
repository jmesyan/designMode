package main

import (
	"net"
	"sync"

	proto "github.com/golang/protobuf/proto"
)

type tcpSession struct {
	conn       *net.TCPConn
	readBuffer tcpReadBuffer
	readChan   chan tcpReadMessage
	sendChan   chan tcpServerMessage
	wait       sync.WaitGroup
}

func newTCPSession(conn *net.TCPConn, readChan chan tcpReadMessage) *tcpSession {
	session := tcpSession{
		conn:       conn,
		readChan:   readChan,
		sendChan:   make(chan tcpServerMessage),
		readBuffer: newTCPReadBuffer(),
	}
	conn.SetReadBuffer(1024 * 32)
	conn.SetWriteBuffer(1024 * 32)
	session.wait.Add(2)
	go session.readLoop()
	go session.writeLoop()
	go session.closeLoop()
	spruceLog.info("accept client from %s ...", (*conn).RemoteAddr().String())
	return &session
}

func (s *tcpSession) close() {
	(*s.conn).Close()
}

func (s *tcpSession) error() {

}

func (s *tcpSession) readLoop() {
	for {
		var err error
		err = s.readBuffer.read(s.conn)
		if nil != err {
			s.sendCloseMsg()
			break
		}
		var hasBody bool
		for {
			hasBody, err = s.readBuffer.readHeader()
			if nil != err {
				break
			}
			if hasBody {
				s.readChan <- s.readBuffer.readMessage(s.conn)
			} else {
				break
			}
		}
		if nil != err {
			spruceLog.error("tcp read error:%s", err.Error())
			s.sendCloseMsg()
			break
		}
	}
	s.wait.Done()
}

func (s *tcpSession) writeLoop() {
	for {
		select {
		case writeMessage := <-s.sendChan:
			if writeMessageFlag == writeMessage.flag {
				(*s.conn).Write(writeMessage.context)
			} else if stopFlag == writeMessage.flag {
				goto stopWrite
			}
		}
	}
stopWrite:
	s.wait.Done()
}

func (s *tcpSession) closeLoop() {
	s.wait.Wait()
	close(s.sendChan)
	spruceLog.info("closed session from client %s ...", (*s.conn).RemoteAddr().String())
}

func (s *tcpSession) sendCloseMsg() {
	closeMessage := tcpServerMessage{flag: stopFlag}
	s.sendChan <- closeMessage
}

func (s *tcpSession) sendMessage(header *MessageHeader, context []byte) {
	headerBytes, err := proto.Marshal(header)
	if nil != err {
		return
	}
	sumLen := len(headerBytes) + len(context)
	messageContext := make([]byte, sumLen)
	copy(messageContext, headerBytes)
	copy(messageContext[len(headerBytes):], context)
	sendMessage := tcpServerMessage{
		flag:    writeMessageFlag,
		context: messageContext,
	}
	s.sendChan <- sendMessage
}
