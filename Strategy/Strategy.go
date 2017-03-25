package main

import "fmt"

type Strategy interface {
	doOperation(num1,num2 int) int
}

type OperationAdd struct {

}

func(o *OperationAdd)doOperation(num1,num2 int) int{
	return num1+num2
}

type OperationSubstract struct {

}

func(o *OperationSubstract)doOperation(num1,num2 int) int{
	return num1-num2
}


type OperationMultiply struct {

}

func(o *OperationMultiply)doOperation(num1,num2 int) int{
	return num1*num2
}


type Context struct {
	strategy Strategy
}

func(c *Context)executeStrategy(num1,num2 int)int{
	return c.strategy.doOperation(num1,num2)
}


func main(){
	context:=&Context{new(OperationAdd)}
	fmt.Println("10+5=",context.executeStrategy(10,5))

	context=&Context{new(OperationSubstract)}
	fmt.Println("10-5=",context.executeStrategy(10,5))

	context=&Context{new(OperationMultiply)}
	fmt.Println("10*5=",context.executeStrategy(10,5))

}