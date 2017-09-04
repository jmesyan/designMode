package main

import (
	"errors"
	"net"

	proto "github.com/golang/protobuf/proto"
)

type tcpReadBuffer struct {
	data       []byte
	dataLen    uint32
	offset     uint32
	readset    uint32
	header     MessageHeader
	headerFlag bool
}

func newTCPReadBuffer() tcpReadBuffer {
	return tcpReadBuffer{
		data:       make([]byte, 1024*64),
		dataLen:    1024 * 64,
		offset:     0,
		readset:    0,
		header:     MessageHeader{},
		headerFlag: false,
	}
}

func (t *tcpReadBuffer) read(conn *net.TCPConn) error {
	offSetData := t.data[t.offset:t.dataLen]
	readLen, err := conn.Read(offSetData)
	t.offset += (uint32)(readLen)
	return err
}

func (t *tcpReadBuffer) readHeader() (bool, error) {
	//没有读取过包头
	if t.headerFlag == false {
		if 6 > (t.offset - t.readset) {
			t.resetBuffer()
			return false, nil
		}
		headData := t.data[t.readset : t.readset+6]
		err := proto.Unmarshal(headData, &t.header)
		if nil != err {
			return false, err
		}
		t.readset += 6
		t.headerFlag = true
		if t.header.BodyLength > t.dataLen {
			return false, errors.New("msg body is to long")
		}
		if t.header.BodyLength > (t.offset - t.readset) {
			return false, nil
		}
	}
	return true, nil
}

func (t *tcpReadBuffer) readMessage(conn *net.TCPConn) tcpReadMessage {
	oldset := t.readset
	t.readset += t.header.BodyLength
	t.headerFlag = false
	bodybytes := make([]byte, t.readset-oldset)
	copy(bodybytes, t.data[oldset:t.readset])
	return tcpReadMessage{
		remoteIP: conn.RemoteAddr().String(),
		localIP:  conn.LocalAddr().String(),
		header:   t.header,
		body:     bodybytes,
	}
}

func moveBytes2Top(bytes []byte, begin uint32, end uint32) {
	copy(bytes[:end-begin+1], bytes[begin:end+1])
}

func (t *tcpReadBuffer) resetBuffer() {
	moveBytes2Top(t.data, t.readset, t.offset)
	t.offset = t.offset - t.readset
	t.readset = 0
}
