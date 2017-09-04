package main

import (
	"time"
)

type moduleTest struct {
	ModuleBase
}

func (m *moduleTest) initModule() {
	m.resgiter(2, handleHello)
}

func handleHello(message *tcpReadMessage) {
	spruceLog.info("%s", message.body)
}

func initModule() {
	spruceModule.resgiter(1, &moduleTest{})
}

func onTick(t time.Time) {
	//spruceLog.info("now: %d", t.UnixNano()/1000000000)
}
func main() {
	spruceIns.init(onTick)
	spruceIns.tickTime = time.Second * 5
	initModule()
	spruceIns.tcpServer(50101)
	cSelf := spruceIns.tcpConnect("self", "", 50101)
	test := MessageHeader{
		BodyLength: 5,
		ModuleId:   1,
		CallBackId: 2,
		ServerInfo: false,
	}
	context := []byte{'h', 'e', 'l', 'l', 'o'}
	cSelf.session.sendMessage(&test, context)
	spruceIns.run()
}
