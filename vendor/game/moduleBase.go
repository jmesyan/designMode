package main

type iModule interface {
	initModule()
}

type moduleInfo struct {
	base   *ModuleBase
	module iModule
}

type messageCb func(message *tcpReadMessage)

type ModuleBase struct {
	id        uint32
	name      string
	callbacks map[uint32]messageCb
}

func (m *ModuleBase) initBase(id uint32, name string) {
	m.id = id
	m.name = name
	m.callbacks = make(map[uint32]messageCb)
}

func (m *ModuleBase) resgiter(id uint32, cb messageCb) {
	m.callbacks[id] = cb
}

func (m *ModuleBase) dispatch(message *tcpReadMessage) {
	callBack, exist := m.callbacks[message.header.CallBackId]
	if exist {
		callBack(message)
	}
}
