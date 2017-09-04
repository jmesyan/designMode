package main

import (
	"bytes"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

const (
	logLevelError = iota
	logLevelWarning
	logLevelInfo
	logNoFile
	logShortFile
	logLongFile
)

var spruceLog iLogger
var spruceLogSingle *logger

type iLogger interface {
	error(format string, v ...interface{})
	warning(format string, v ...interface{})
	info(format string, v ...interface{})
}

type logger struct {
	logLevel      int
	logFlag       int
	errorLogger   *log.Logger
	warningLogger *log.Logger
	infoLogger    *log.Logger
	consoleLogger *log.Logger
}

func logInit(debug bool, level int, flag int, logRoot string) {
	if nil == spruceLogSingle {
		if true == debug {
			spruceLogSingle = &logger{
				logLevel:      level,
				logFlag:       flag,
				consoleLogger: log.New(os.Stdout, "", 0),
			}
		} else {
			fileNameFront := logRoot + "_" + logGetDate()
			spruceLogSingle = &logger{
				logLevel:      level,
				errorLogger:   logCreateFile(fileNameFront + "_error.log"),
				warningLogger: logCreateFile(fileNameFront + "_warning.log"),
				infoLogger:    logCreateFile(fileNameFront + "_info.log"),
			}
		}
		spruceLog = spruceLogSingle
	}
}

func (l *logger) error(format string, v ...interface{}) {
	if nil == l || logLevelError > l.logLevel {
		return
	}
	var buffer bytes.Buffer
	if nil != l.errorLogger {
		l.logGetHeader(&buffer, 3, logLevelError, false)
		buffer.WriteString(format)
		l.errorLogger.Printf(buffer.String(), v...)
	}
	if nil != l.consoleLogger {
		l.logGetHeader(&buffer, 3, logLevelError, true)
		buffer.WriteString(format)
		l.consoleLogger.Printf(buffer.String(), v...)
	}
}

func (l *logger) warning(format string, v ...interface{}) {
	if nil == l || logLevelWarning > l.logLevel {
		return
	}
	var buffer bytes.Buffer
	if nil != l.warningLogger {
		l.logGetHeader(&buffer, 3, logLevelWarning, false)
		buffer.WriteString(format)
		l.warningLogger.Printf(buffer.String(), v...)
	}
	if nil != l.consoleLogger {
		l.logGetHeader(&buffer, 3, logLevelWarning, true)
		buffer.WriteString(format)
		l.consoleLogger.Printf(buffer.String(), v...)
	}
}

func (l *logger) info(format string, v ...interface{}) {
	if nil == l || logLevelInfo > l.logLevel {
		return
	}
	var buffer bytes.Buffer
	if nil != l.infoLogger {
		l.logGetHeader(&buffer, 3, logLevelInfo, false)
		buffer.WriteString(format)
		l.infoLogger.Printf(buffer.String(), v...)
	}
	if nil != l.consoleLogger {
		l.logGetHeader(&buffer, 3, logLevelInfo, true)
		buffer.WriteString(format)
		l.consoleLogger.Printf(buffer.String(), v...)
	}
}

func logGetDate() string {
	now := time.Now()
	var timeBuffer bytes.Buffer
	timeBuffer.WriteString(strconv.Itoa(now.Year()))
	timeBuffer.WriteString("_")
	timeBuffer.WriteString(strconv.Itoa(int(now.Month())))
	timeBuffer.WriteString("_")
	timeBuffer.WriteString(strconv.Itoa(now.Day()))
	timeBuffer.WriteString("_")
	timeBuffer.WriteString(strconv.Itoa(now.Hour()))
	return timeBuffer.String()
}

func logCreateFile(path string) *log.Logger {
	var file *os.File
	var err error
	if _, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(path)
		}
	} else {
		file, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	}
	if err != nil {
		log.Printf("create or open %s failed! %s \n", path, err)
	}
	return log.New(file, "", 0)
}

func (l *logger) logGetHeader(buffer *bytes.Buffer, calldepth int, level int, needColour bool) {
	//write level
	l.logGetHeaderLevel(buffer, level, needColour)
	//wtite time
	l.logGetHeaderTime(buffer)
	//write file
	l.logGetHeaderFile(buffer, calldepth)
}

func (l *logger) logGetHeaderLevel(buffer *bytes.Buffer, level int, needColour bool) {
	switch level {
	case logLevelError:
		if true == needColour {
			buffer.WriteString("\x1b[91m[spruce error]")
		} else {
			buffer.WriteString("[spruce error]")
		}
	case logLevelWarning:
		if true == needColour {
			buffer.WriteString("\x1b[93m[spruce warning]")
		} else {
			buffer.WriteString("[spruce warning]")
		}
	case logLevelInfo:
		buffer.WriteString("[spruce info]")
	}
}

func (l *logger) logGetHeaderTime(buffer *bytes.Buffer) {
	now := time.Now() // get this early.
	buffer.WriteString("[")
	buffer.WriteString(strconv.Itoa(now.Hour()))
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(now.Minute()))
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(now.Second()))
	buffer.WriteString("]")
}

func (l *logger) logGetHeaderFile(buffer *bytes.Buffer, calldepth int) {
	if logShortFile != l.logFlag && logLongFile != l.logFlag {
		return
	}
	var file string
	var line int
	var ok bool
	_, file, line, ok = runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	if l.logFlag == logShortFile {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if (file)[i] == '/' {
				short = (file)[i+1:]
				break
			}
		}
		file = short
	}
	buffer.WriteString("[")
	buffer.WriteString(file)
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(line))
	buffer.WriteString("]")
}
