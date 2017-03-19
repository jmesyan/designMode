package main

import "fmt"

type DrawAPI interface {
	drawCircle(radius,x,y int)
}


type RedCircle struct {

}

func(r *RedCircle)drawCircle(radius,x,y int){
	fmt.Println("Drawing Circle[ color: red, radius: ",radius, ", x: ",x,", y:",y,"]\r\n")
}

type GreenCircle struct {

}

func (g *GreenCircle)drawCircle(radius,x,y int){
	fmt.Println("Drawing Circle[ color: green, radius: ",radius, ", x: ",x,", y:",y,"]\r\n")
}

type InterShape interface {
	draw()
}


type Shape struct {
	drawAPI DrawAPI
}

func(s *Shape)init(drawAPI DrawAPI){
	s.drawAPI=drawAPI
}


type Circle struct {
	Shape
	x int
	y int
	radius int
}

func(c *Circle)init(x,y,radius int,drawAPI DrawAPI){
	c.Shape.init(drawAPI)
	c.x=x
	c.y=y
	c.radius=radius
}

func(c *Circle)draw(){
	c.drawAPI.drawCircle(c.radius,c.x,c.y)
}


func main(){
	redCircle:=new(Circle)
	redCircle.init(100,100,10,new(RedCircle))
	greenCircle:=new(Circle)
	greenCircle.init(150,150,15,new(GreenCircle))
	redCircle.draw()
	greenCircle.draw()
}


