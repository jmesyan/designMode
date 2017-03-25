package main

import (
	"container/list"
	"fmt"
)

type Observer interface {
	update()
}

type Subject struct {
	state int
	observers *list.List
}

func(s *Subject)getState() int{
	return s.state
}

func(s *Subject)setState(state int){
	s.state=state
	s.notifyObservers()
}

func(s *Subject)attach(observer Observer){
	s.observers.PushBack(observer)
}

func(s *Subject)notifyObservers(){
	for o:=s.observers.Front();o!=nil;o=o.Next(){
		observer:=o.Value.(Observer)
		observer.update()
	}
}


type BinaryObserver struct {
	subject *Subject
}

func (b *BinaryObserver)init(subject *Subject){
	b.subject=subject
	b.subject.attach(b)
}

func(b *BinaryObserver)update(){
	fmt.Println("this is binary string",b.subject.getState())
}

type OctalObserver struct {
	subject *Subject
}

func (b *OctalObserver)init(subject *Subject){
	b.subject=subject
	b.subject.attach(b)
}

func(b *OctalObserver)update(){
	fmt.Println("this is octal string",b.subject.getState())
}


type HexaObserver struct {
	subject *Subject
}

func (b *HexaObserver)init(subject *Subject){
	b.subject=subject
	b.subject.attach(b)
}

func(b *HexaObserver)update(){
	fmt.Println("this is hexa string",b.subject.getState())
}


func main(){
	subject:=&Subject{observers:list.New()};
	o1:=new(BinaryObserver)
	o1.init(subject)
	o2:=new(OctalObserver)
	o2.init(subject)
	o3:=new(HexaObserver)
	o3.init(subject)

	fmt.Println("first state changes:15")
	subject.setState(15)

	fmt.Println("secend state changes:10")
	subject.setState(10)
}