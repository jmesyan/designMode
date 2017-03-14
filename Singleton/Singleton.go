package main

import "fmt"

var _instance *Singleton

type Singleton struct {
	name string
}

func getInstance() *Singleton {
	if _instance == nil {
		_instance = new(Singleton)
		fmt.Println(11111)
	}
	fmt.Println(22222);
	return _instance
}

func (this *Singleton) setName(name string) {
	this.name = name;
}

func (this *Singleton) showMessage() {
	fmt.Println("hello world", this.name);
}

func main() {
	object := getInstance();
	object.setName("object1")
	object.showMessage();
	object2 := getInstance();
	object2.setName("object2")
	object2.showMessage();
}