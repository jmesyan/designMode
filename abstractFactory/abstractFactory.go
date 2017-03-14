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

//color接口
type color interface {
	fill()
}

//红色
type Red struct {

}

func (color *Red)fill() {
	fmt.Print("Red fill() method\r\n");
}


//绿色
type Green struct {

}

func (color *Green)fill() {
	fmt.Print("Green fill() method\r\n");
}

//蓝色
type Blue struct {

}

func (color *Blue)fill() {
	fmt.Print("Blue fill() method\r\n");
}

type AbstructFactory interface {
	getShape(string) shape
	getColor(string) color
}

type shapeFactory struct {

}

func (AbstructFactory *shapeFactory) getShape(shapeType string) (shapeResult shape) {
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

func (AbstructFactory *shapeFactory)getColor(color string)color{
	return nil
}

type colorFactory struct {

}

func (AbstructFactory *colorFactory) getColor(color string) (colorResult color) {
	if (color == "") {
		return nil
	}
	color = strings.ToLower(color)
	switch color {
	case "red":
		colorResult = new(Red)
	case "green":
		colorResult = new(Green)
	case "blue":
		colorResult = new(Blue)
	default:
		panic("no color find!")

	}
	return
}

func (AbstructFactory *colorFactory)getShape(color string)shape{
	return nil
}

type FactoryProducer struct {

}

func(this *FactoryProducer)getFactory(choice string) AbstructFactory{
	choice=strings.ToLower(choice)
	if(choice=="shape"){
		return new(shapeFactory)
	}else if(choice=="color"){
		return new(colorFactory)
	}
	return nil
}

func main(){
	//demo
	fp:=new(FactoryProducer)
	sfp:=fp.getFactory("shape")
	shape1:=sfp.getShape("rectangle")
	shape1.draw()
	shape2:=sfp.getShape("circle")
	shape2.draw()

	cfp:=fp.getFactory("color")
	color1:=cfp.getColor("Green")
	color1.fill()
	color2:=cfp.getColor("Red")
	color2.fill()

}


