package main

import "fmt"

type Context struct {
     state State
}

func(c *Context)setState(state State){
	c.state=state
}

func(c *Context)getState() State{
	return c.state
}


type State interface {
	doAction(context *Context)
	toString()
}


type StartState struct {

}

func(s *StartState)doAction(context *Context){
	fmt.Println("player is in start state")
	context.setState(s)
}

func(s *StartState)toString(){
	fmt.Println("Start State")
}

type StopState struct {

}

func(s *StopState)doAction(context *Context){
	fmt.Println("player is in stop state")
	context.setState(s)
}

func(s *StopState)toString(){
	fmt.Println("Stop state")
}


func main(){
	context:=new(Context)
	startState:=new(StartState)
	startState.doAction(context)
	context.getState().toString()
	stopState:=new(StopState)
	stopState.doAction(context)
	context.getState().toString()
}
