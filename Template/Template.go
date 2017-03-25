package main

import "fmt"

type Game interface {
	initialize();
	startPlay();
	endPlay();
}

func play(g Game){
	g.initialize()
	g.startPlay()
	g.endPlay()
}


type Cricket struct {

}

func(c *Cricket)initialize(){
	fmt.Println("Cricket Game Initialized! Start playing.");
}

func(c *Cricket)startPlay(){
	fmt.Println("Cricket Game Started. Enjoy the game!");
}

func(c *Cricket)endPlay(){
	fmt.Println("Cricket Game Finished!");
}

type Football struct {

}

func(c *Football)initialize(){
	fmt.Println("Football Game Initialized! Start playing.");
}

func(c *Football)startPlay(){
	fmt.Println("Football Game Started. Enjoy the game!");
}

func(c *Football)endPlay(){
	fmt.Println("Football Game Finished!");
}

func main(){
	//demo
	cricket:=new(Cricket)
	play(cricket)

	football:=new(Football)
	play(football)
}


