package main

import (
	"fmt"
	"strings"
)

//shape接口
type shape interface {
	draw()
}

//矩形
type Rectangle struct {

}

func (shape *Rectangle)draw() {
	fmt.Print("Rectangle draw() method\r\n");
}


//方形
type Square struct {

}

func (shape *Square)draw() {
	fmt.Print("Square draw() method\r\n");
}

//圆形
type Circle struct {

}

func (shape *Circle)draw() {
	fmt.Print("Circle draw() method\r\n");
}

type shapeFactory struct {

}

func (this *shapeFactory)getShape(shapeType string) (shapeResult shape) {
	if (shapeType == "") {
		return nil
	}
	shapeType = strings.ToLower(shapeType)
	switch shapeType {
	case "rectangle":
		shapeResult = new(Rectangle)
	case "square":
		shapeResult = new(Square)
	case "circle":
		shapeResult = new(Circle)
	default:
		panic("no graph find!")

	}
	return
}


func main(){
	defer func(){
	      if err:=recover();err!=nil{
		      fmt.Println(err)
	      }
	}()

	var fac shapeFactory
	shape:=fac.getShape("circle");
	shape.draw()
	shape=fac.getShape("rectangle");
	shape.draw()
	shape=fac.getShape("abc");
	shape.draw()
}








