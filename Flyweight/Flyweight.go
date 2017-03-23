package main

import (
	"fmt"
	"math/rand"
)

type Shape interface {
	draw()
}

type Circle struct {
	color  string
	x      int
	y      int
	radius int
}

func (c *Circle)setX(x int) {
	c.x = x
}

func (c *Circle)setY(y int) {
	c.y = y
}

func (c *Circle)setRadius(radius int) {
	c.radius = radius
}

func (c *Circle)draw() {
	fmt.Println("Circle: Draw() [Color : ", c.color, ", x : ", c.x, ", y :", c.y, ", radius :", c.radius)
}

type ShapeFactory struct {
	shapeMap map[string]*Circle
}

func(s *ShapeFactory)init(){
	s.shapeMap=make(map[string]*Circle)
}

func (s *ShapeFactory)getCircle(color string) *Circle {
	shape := s.shapeMap[color];
	if shape == nil{
		shape = &Circle{color:color}
		s.shapeMap[color]=shape
		fmt.Println("Creating circle of color : ",color)
	}
	return shape
}

func main() {
	colors := []string{"Red", "Green", "Blue", "White", "Black"}
	shapeFactory := new(ShapeFactory)
	shapeFactory.init()
	for i := 1; i <= 20; i++ {
		cn := rand.Intn(5);
		circle := shapeFactory.getCircle(colors[cn])
		circle.setX(rand.Intn(100))
		circle.setY(rand.Intn(100))
		circle.setRadius(100)
		circle.draw()

	}

}

