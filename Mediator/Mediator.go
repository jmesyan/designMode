package main

import (
	"fmt"
	"time"
)
type ChatRoom struct {

}
func(c *ChatRoom)sendMessage(user *User,message string){
	fmt.Println(time.Now().Format("2016-01-02 15:04:05"),":[",user.getName(),"]-",message)
}
var chatRoom ChatRoom;

type User struct {
	name string
}
func(u *User)getName()string{
	return u.name;
}

func(u *User)setName(name string){
	u.name=name;
}

func(u *User)sendMessage(message string){
      chatRoom.sendMessage(u,message)
}


func main(){
	robert:=&User{"robert"}
	john:=&User{"John"}
	robert.sendMessage("hi John")
	john.sendMessage("hello robert");
}
