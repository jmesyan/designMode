package main

import "fmt"

//shape接口
type Shape interface {
	draw()
}

//矩形
type Rectangle struct {

}

func (s *Rectangle)draw() {
	fmt.Print("Rectangle draw() method\r\n");
}


//方形
type Square struct {

}

func (s *Square)draw() {
	fmt.Print("Square draw() method\r\n");
}

//圆形
type Circle struct {

}

func (s *Circle)draw() {
	fmt.Print("Circle draw() method\r\n");
}

type ShapeMaker struct {
	circle Shape
	rectangle Shape
	square Shape
}

func(s *ShapeMaker)drawCircle(){
	s.circle.draw()
}

func(s *ShapeMaker)drawRectangle(){
	s.rectangle.draw()
}

func(s *ShapeMaker)drawSquare(){
	s.square.draw()
}


func main(){
	shapeMaker:=&ShapeMaker{new(Circle),new(Rectangle),new(Square)}
	shapeMaker.drawCircle()
	shapeMaker.drawRectangle()
	shapeMaker.drawSquare()
}




