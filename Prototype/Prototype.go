package main

import (
   "fmt"
)

type InterShape interface {
   getId() int
   getType() string
   setId(int)
   setType()
   setTest(string)
   draw()
}

//父类
type Shape struct {
    id int
    Type string
}

func(s *Shape)setTest(name string){
   s.Type=name
}

func(s *Shape) getType() string{
   return s.Type
}

func(s *Shape) getId() int{
   return s.id
}

func(s *Shape)setId(id int){
   s.id=id
}

//子类
type Rectangle struct {
   Shape
}

func(s *Rectangle)setType(){
   s.Type="Rectangle"
}


func(s *Rectangle)draw(){
   fmt.Println("Inside Rectangle::draw() method");
}

type Circle struct {
   Shape
}

func(s *Circle)setType(){
   s.Type="Circle"
}


func(s *Circle)draw(){
   fmt.Println("Inside Circle::draw() method");
}

type Square struct {
   Shape
}

func(s *Square)setType(){
   s.Type="Square"
}


func(s *Square)draw(){
   fmt.Println("Inside Square::draw() method");
}

//创建shapeCache
type ShapeCache struct {
   shapeMap map[int]InterShape
}

func(s *ShapeCache)getShape(shapeId int) InterShape{
    shape:=s.shapeMap[shapeId]
    return shape//如何实现对象赋值？
}

func(s *ShapeCache)loadCache(){
   s.shapeMap=make(map[int]InterShape)
   circle:=new(Circle)
   circle.setId(1)
   circle.setType()
   s.shapeMap[circle.getId()]=circle

   square:=new(Square)
   square.setId(2)
   square.setType()
   s.shapeMap[square.getId()]=square

   rectangle:=new(Rectangle)
   rectangle.setId(3)
   rectangle.setType()
   s.shapeMap[rectangle.getId()]=rectangle
}


func main(){
   shapeCache:=new(ShapeCache)
   shapeCache.loadCache()
   clonedShape1:=shapeCache.getShape(1)
   fmt.Println("Shape1:",clonedShape1.getType())
   clonedShape1.draw()

   clonedShape2:=shapeCache.getShape(2)
   fmt.Println("Shape2:",clonedShape2.getType())
   clonedShape2.draw()

   clonedShape3:=shapeCache.getShape(3)
   fmt.Println("Shape3:",clonedShape3.getType())
   clonedShape3.draw()

   clonedShape4:=shapeCache.getShape(1)
   clonedShape4.setTest("Hello")
   fmt.Println("Shape4:",clonedShape4.getType())
   clonedShape4.draw()
   fmt.Println("Shape1:",clonedShape1.getType())
}


