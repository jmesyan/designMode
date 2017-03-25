package main

import "fmt"

type Itertor interface {
	init(names []string,index int)
	hasNext() bool
	next() string
}

type NamesIteror struct {
	names []string
	index int
}

func(n *NamesIteror)init(names []string,index int){
	n.names=names
	n.index=index
}

func(n *NamesIteror)hasNext()bool{
	if(n.index<len(n.names)){
		return true;
	}
	return false;
}

func(n *NamesIteror)next() string{
	if(n.hasNext()){
		index:=n.index
		n.index++
		return n.names[index]
	}
	return "";
}


type Container interface {
	getItertor() Itertor
}

type NameRepository struct {
	names []string
	itertor Itertor
}

func(n *NameRepository)getItertor() Itertor{
	n.names=[]string{0:"Robert",1:"John",2:"Julie",3:"Lora"}
	n.itertor=new(NamesIteror)
	n.itertor.init(n.names,0)
	return n.itertor
}

func main(){
	//demo
	nameRepository:=new(NameRepository)
	nameR:=nameRepository.getItertor()
	for nr:=nameR;nr.hasNext();{
		fmt.Println("Name :",nr.next())
	}
}