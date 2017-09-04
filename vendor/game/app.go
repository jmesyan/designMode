package main

import (
	"strconv"
	"time"
)

type iApp interface {
	init()
}

type app struct {
	startParm parms
	readChan  chan tcpReadMessage
	tickTime  time.Duration
	cbOnTick  func(time.Time)
	connects  map[string]iConnect
}

var spruceIns *app

func init() {
	spruceIns = new(app)
	spruceIns.connects = make(map[string]iConnect)
	spruceIns.tickTime = time.Second * 2
}

func (a *app) init(cbOnTick func(time.Time)) {
	a.startParm = parseParms()
	logInit(a.startParm.logDebug, a.startParm.logLevel,
		a.startParm.logFlag, a.startParm.logRoot)
	a.readChan = make(chan tcpReadMessage, 1024)
	a.cbOnTick = cbOnTick
}

//之後可能改成多個端口
func (a *app) tcpServer(port int) *tcpServer {
	return newTCPServer(port, a.readChan)
}

func (a *app) tcpConnect(name string, ip string, port int) *tcpConnect {
	addr := ip + ":" + strconv.Itoa(port)
	connect := newTCPConnect(addr, a.readChan)
	a.connects[name] = connect
	return connect
}

func (a *app) run() {
	for {
		select {
		case message := <-a.readChan:
			spruceModule.dispatch(&message)
		case <-time.After(a.tickTime):
			if nil != a.cbOnTick {
				a.cbOnTick(time.Now())
			}

		}
	}

}
