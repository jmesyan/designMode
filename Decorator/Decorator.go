package main

import "fmt"

type Shape interface {
	draw()
}

type Circle struct {

}
func(c *Circle)draw(){
	fmt.Println("shape:Circle")
}


type Rectangle struct {

}

func(r *Rectangle)draw(){
	fmt.Println("shape:Rectangle")
}

type ShapeDector struct {
	dectorShape Shape
}

func(s *ShapeDector)draw(){
	s.dectorShape.draw()
}

type RedShapeDector struct {
	*ShapeDector
}
func(r *RedShapeDector)setRedBorder(shape Shape){
   fmt.Println("Bolder Color:Red")
}

func(r *RedShapeDector)draw(){
	r.ShapeDector.draw()
	r.setRedBorder(r.dectorShape)
}

func main(){
	circle:=new(Circle)
	redCircle:=&RedShapeDector{new(Circle)}
	redRectangle:=&RedShapeDector{new(Rectangle)}

	circle.draw()
	redCircle.draw()
	redRectangle.draw()
}