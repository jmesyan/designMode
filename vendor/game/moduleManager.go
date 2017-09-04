package main

import (
	"reflect"
	"unsafe"
)

type moduleManager struct {
	modules map[uint32]moduleInfo
}

var spruceModule *moduleManager

func init() {
	spruceModule = newModuleManager()
}

func newModuleManager() *moduleManager {
	return &moduleManager{modules: make(map[uint32]moduleInfo)}
}

func (m *moduleManager) resgiter(id uint32, module iModule) {
	moduleType := reflect.TypeOf(module).Elem()
	moduleName := moduleType.Name()
	if _, exist := moduleType.FieldByName("ModuleBase"); false == exist {
		spruceLog.error("%s id %d is not inherit from ModuleBase!", moduleName, id)
		return
	}
	_, exist := m.modules[id]
	if exist {
		spruceLog.error("%s id %d is exist!", moduleName, id)
		return
	}
	moduleVal := reflect.ValueOf(module).Elem()
	moduleBaseValue := moduleVal.FieldByName("ModuleBase")
	moduleBaseUnsafePtr := unsafe.Pointer(moduleBaseValue.UnsafeAddr())
	moduleBasePtr := (*ModuleBase)(moduleBaseUnsafePtr)
	m.modules[id] = moduleInfo{
		base:   moduleBasePtr,
		module: module,
	}
	m.modules[id].base.initBase(id, moduleName)
	m.modules[id].module.initModule()
	spruceLog.info("%s id %d is resgiter successd!", moduleName, id)
}

func (m *moduleManager) dispatch(message *tcpReadMessage) {
	module, exist := m.modules[message.header.ModuleId]
	if exist {
		module.base.dispatch(message)
	}
}
