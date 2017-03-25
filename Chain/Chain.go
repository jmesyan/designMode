package main

import "fmt"

const (
	DEBUG = 3
	Error = 2
	INFO = 1
)
type Logger interface {
	write(message string)
	getLevel()int
	getNextLogger() Logger

}

type AbstractLogger struct {
	nextLogger Logger
	level      int
}
func(a *AbstractLogger)setNextLogger(nextLogger Logger){
   a.nextLogger=nextLogger
}

func (a *AbstractLogger)getLevel() int{
	return a.level
}

func (a *AbstractLogger)getNextLogger() Logger{
	return a.nextLogger
}

func LogMessage(a Logger,level int,message string){
	if(a.getLevel()<=level){
		a.write(message)
	}

	if(a.getNextLogger()!=nil){
		LogMessage(a.getNextLogger(),level,message)
	}
}

type ConsoleLogger struct {
	AbstractLogger
}

func(e *ConsoleLogger)init(level int){
	e.level=level;
}


func(e *ConsoleLogger)write(message string){
	fmt.Println("Standard Console::Logger: ",message)
}

type ErrorLogger struct {
	AbstractLogger
}

func(e *ErrorLogger)init(level int){
	e.level=level;
}


func(e *ErrorLogger)write(message string){
    fmt.Println("Error Console::Logger: ",message)
}

type FileLogger struct {
	AbstractLogger
}

func(e *FileLogger)init(level int){
	e.level=level;
}

func(e *FileLogger)write(message string){
	fmt.Println("File::Logger: ",message)
}


func chainPatternDemo() Logger{
	errorLogger:=new(ErrorLogger)
	errorLogger.init(Error)
	fileLogger:=new(FileLogger)
	fileLogger.init(DEBUG)
	consoleLogger:=new(ConsoleLogger)
	consoleLogger.init(INFO)
	errorLogger.setNextLogger(fileLogger)
	fileLogger.setNextLogger(consoleLogger)
	return errorLogger
}

func main() {
	logChain:=chainPatternDemo()
	LogMessage(logChain,INFO,"this is an INFO info");
	LogMessage(logChain,DEBUG,"this is an DEBUG info");
	LogMessage(logChain,Error,"this is an Error info");
}
