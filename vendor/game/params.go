package main

import (
	"flag"
	"runtime"
)

const (
	Linux = iota
	Window
	DebugMode
	ReleaseMode
)

type parms struct {
	platform int
	mode     int
	logDebug bool
	logLevel int
	logFlag  int
	logRoot  string
}

func parseParms() parms {
	p := parms{}
	p.logDebug = *flag.Bool("mode", true, "debug")
	if "windows" == runtime.GOOS {
		p.platform = Window
		p.logRoot = "..\\log"
	} else if "linux" == runtime.GOOS {
		p.platform = Linux
		p.logRoot = "..\\log"
	} else {
		p.platform = Linux
		p.logRoot = "..\\log"
	}

	if true == p.logDebug {
		p.mode = DebugMode
		p.logDebug = true
		p.logLevel = logLevelInfo
		p.logFlag = logLongFile
	} else { //release
		p.mode = ReleaseMode
		p.logDebug = *flag.Bool("logDebug", false, "log open console")
		p.logLevel = *flag.Int("logLevel", logLevelError, "log level")
		p.logFlag = *flag.Int("logFlag", logNoFile, "log header function tips")
	}
	return p
}
